{{ template "base.tpl" . }}

{{ define "content" }}
  <article>
    {{if .notfound }}
      <h1>Not found</h1>
	{{else}}
      <h1>{{.post.Title}}</h1>
      <div class="content">
	    {{.post.Content}}
      </div>
      <div class="footer border-bottom text-right">
        {{.post.Member.Username}} at {{date .post.PostedAt "Y-m-d h:i:s"}}
      </div>
	{{end}}
{{ end }}

{{ define "scripts" }}
{{ end }}
