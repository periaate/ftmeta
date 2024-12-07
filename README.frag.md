# {{ .Module }}
{{ .Desc }}

## Demonstration
We will be importing the `desc.frag.md` fragment from another repository, {{ link "blume" }}, which I have locally stored.

### {{ link "blume" }}
{{ desc "blume" }}
{{ import "blume" "sub" }}

This was templated with the following code from {{ link "blume" }}:
```
### {{ `{{` }} link "blume" {{ `}}` }}
{{ `{{` }} desc "blume" {{ `}}` }}
{{ `{{` }} import "blume" "sub" {{ `}}` }}
```

Which then generateds the following markdown:
```markdown
### [blume](https://github.com/periaate/blume)
blume is a programming language, embedded into Go in the form of a standard library. It wraps around Go's existing standard libraries, or implements them from scratch, with internally consistent semantics.
- gen consists of generic functions, types, etc.
- yap is a much simpler `log/slog` like logger.
```

## Usage
> [!CAUTION]
> Currently ftmeta is early in development, and is being actively developed with hard coded values. As such, without changes ftmeta can not be used outside my environment.

