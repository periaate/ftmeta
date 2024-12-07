package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/periaate/blume/fsio"
	. "github.com/periaate/blume/fsio"
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
	"github.com/periaate/blume/hnet"
	"github.com/periaate/blume/lazy"
	"github.com/periaate/blume/typ"
	"github.com/periaate/blume/yap"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"Replace": func(args ...any) typ.String {
			last := len(args) - 1
			var s typ.String
			switch args[last].(type) {
			case typ.String:
				s = args[last].(typ.String)
			case string:
				s = typ.String(args[last].(string))
			default:
				yap.Fatal("last argument must be a string or typ.String")
			}
			sar := make([]string, 0, last)
			for i := range last {
				sar = append(sar, args[i].(string))
			}
			return typ.String(s).Replace(sar...)
		},
		"Dir":  func(a any) FilePath { return FilePath(a.(string)).Dir() },
		"Base": func(a any) FilePath { return FilePath(a.(string)).Base() },
		"Abs":  func(a any) FilePath { return FilePath(a.(string)).Abs().Unwrap() },
		"Link": func(a any) string { return Link(a.(string)) },
	}
}

var readDir = lazy.Monadic(fsio.ReadsDir[string])

func ReadFrags(dir string) gen.Array[Fragment] {
	res := readDir(dir).Err(func(err T.Error[any]) { yap.Fatal("could't read dir", err) })
	yap.Debug("read fragments", "dir", dir, "res", res)
	res = res.Filter(gen.Contains(fsio.FilePath(".frag.")))
	yap.Debug("read fragments", "dir", dir, "res", res)
	return gen.ArrayFrom(res, func(a fsio.FilePath) Fragment { return Fragment(a.Clean()) })
}

func main() {
	r := Import("./README")
	yap.Debug("rendering fragment", "README")
	fmt.Print(r)
}

type Fragment string

func (f Fragment) URL() string {
	return hnet.URL(FilePath(f).Abs().Unwrap().Dir().Clean()).ReplacePrefix("C:/", "").AsProtocol(hnet.HTTPS).String()
}

func (f Fragment) Module(a ...string) string {
	fp := fsio.FilePath(f)
	abs := fp.Abs().Unwrap()
	dir := abs.Dir()
	base := dir.Base()
	yap.Debug("getting module", "templating", f)
	yap.Debug("abs", abs)
	yap.Debug("dir", dir)
	yap.Debug("base", base)
	yap.Debug("string", base.String())
	return base.String()
}

func (f Fragment) Name() string {
	return typ.String(fsio.FilePath(f).Base()).Split(".").GetShift().Unwrap().String()
}

func (f Fragment) String() string { return string(f) }

func (f Fragment) Template(w io.Writer) {
	yap.Debug("templating", "template path", f)
	raw := string(fsio.Read(f).Unwrap())
	tmpl := gen.Must(template.New(string(f)).Funcs(FuncMap()).Parse(raw))
	yap.ErrFatal(tmpl.Execute(w, f), "could not execute template")
}

func Import(path string) (res string) {
	yap.Debug("importing fragment", "path", path)
	switch {
	case gen.HasPrefix("github.com")(path):
		path = FilePath("C:/" + path).Abs().Unwrap().String()
	case gen.HasPrefix(".")(path):
	default:
		path = "./" + path
	}

	fragName := FilePath(path).Base()
	yap.Debug("importing fragment", "path", path, "fragment", fragName)

	frags := ReadFrags(FilePath(path).Dir().String())
	frag := frags.First(gen.Contains(Fragment(fragName))).Unwrap()
	yap.Debug("importing", "found fragments", frags.Array(), "matched", frag)
	buf := fsio.B()
	frag.Template(buf)
	return typ.String(buf.String()).ReplaceSuffix("\n", "").String()
}

func Link(path string) (res string) {
	yap.Debug("importing fragment", "path", path)
	switch {
	case gen.HasPrefix("github.com")(path):
		path = FilePath("C:/" + path).Abs().Unwrap().String()
	case gen.HasPrefix(".")(path):
	default:
		path = "./" + path
	}

	fragName := FilePath(path).Base()
	yap.Debug("importing fragment", "path", path, "fragment", fragName)

	frags := ReadFrags(FilePath(path).Dir().String())
	frag := frags.First(gen.Contains(Fragment(fragName))).Unwrap()
	yap.Debug("importing", "found fragments", frags.Array(), "matched", frag)
	return fmt.Sprintf("[%s](%s)", frag.Module(), frag.URL())
}

func (f Fragment) Import(path string) string { return Import(path) }

func (f Fragment) Link() (res string) { return Link(f.String()) }
