# {{ .Link }}
{{ .Import "description" }}

## Demonstration
We will be importing `description` fragments from my other repos on my local disk.

### {{ Link "github.com/periaate/blume/description" }}
{{ .Import "github.com/periaate/blume/description" }} 

This was templated with the following code from {{ Link "github.com/periaate/blume/description" }}:
```
### {{`{{`}} Link "github.com/periaate/blume/description" {{`}}`}}
{{`{{`}} .Import "github.com/periaate/blume/description" {{`}}`}}
```
