{{ template "base.tpl" . }}

{{ define "content" }}
  <h1>Login</h1>
  {{if eq true .has_error}}
  <div class="alert alert-danger" role="alert">
    Error: {{.error}}
  </div>
  {{end}}
  <form action="" method="POST">
  {{.xsrfdata}}
  <div class="form-group">
    <label for="username">Username</label>
    <input type="text" name="username" class="form-control" id="username" placeholder="User name">
  </div>
  <div class="form-group">
    <label for="password">Password</label>
    <input type="password" name="password" class="form-control" id="password" placeholder="Password">
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
  </form>
{{ end }}
