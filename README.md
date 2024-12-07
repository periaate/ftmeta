# [ftmeta](https://github.com/periaate/ftmeta)
ftmeta is a preprocessor for Go's `text/template` templating language, providing imports, further string manipulation features, among other improvements.

## Demonstration
We will be importing `description` fragments from my other repos on my local disk.

### [blume](https://github.com/periaate/blume)
 yap is a much simpler `log/slog` like logger.
 gen consists of generic functions, types, etc. 

This was templated with the following code from [blume](https://github.com/periaate/blume):
```
### {{ Link "github.com/periaate/blume/description" }}
{{ .Import "github.com/periaate/blume/description" }}
```