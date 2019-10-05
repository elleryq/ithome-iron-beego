{{ template "base.tpl" . }}

{{ define "extra_head" }}
{{end}}

{{ define "content" }}
  <h1>Delete post</h1>
  {{if .flash.error}}
  <div class="alert alert-danger" role="alert">
    Error: {{.flash.error}}
  </div>
  {{end}}
  {{ if .notfound }}
  <div class="alert alert-danger" role="alert">
    Not found
  </div>
  <a href="{{urlfor "PostController.GetAll"}}" class="btn btn-primary">Return</a>
  {{ else }}
  <form action="" method="POST">
  {{.xsrfdata}}
  <div class="form-group">
    <input type="hidden" name="id" value="{{.post.Id}}">
    <label for="title">Title</label>
    <input type="text" name="title" class="form-control" id="title" placeholder="Title" value="{{.post.Title}}" disabled>
  </div>
  <div class="form-group">
    <label for="content">Content</label>
    {{str2html .post.Content}}
  </div>
  <a href="{{urlfor "PostController.GetAll"}}" class="btn btn-primary">Return</a>
  <button type="submit" class="btn btn-danger">DELETE</button>
  </form>
  {{end}}
{{ end }}

{{ define "scripts" }}
{{ end }}
