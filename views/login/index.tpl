{{ template "base.tpl" . }}

{{ define "content" }}
  <h1>Login</h1>
  {{if .flash.error }}
  <div class="alert alert-danger" role="alert">
    {{.flash.error}}
  </div>
  {{end}}
  {{if .flash.success }}
  <div class="alert alert-success" role="alert">
    {{.flash.success}}
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
