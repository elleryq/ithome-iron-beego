{{ template "base.tpl" . }}

{{ define "content" }}
  <h1>Logout</h1>
  {{if eq true .has_error}}
  <div class="alert alert-danger" role="alert">
    Error: {{.error}}
  </div>
  {{end}}
  <a href="/login">Login</a>
{{ end }}
