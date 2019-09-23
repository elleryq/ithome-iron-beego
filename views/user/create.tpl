{{ template "base.tpl" . }}

{{ define "content" }}
  <h1>User - Create</h1>
  {{if .flash.error}}
  <div class="alert alert-danger" role="alert">
    Error: {{.flash.error}}
  </div>
  {{end}}
  <form action="" method="POST">
  {{.xsrfdata}}
  <div class="form-group">
    <label for="name">Name</label>
    <input type="text" name="name" class="form-control" id="name" placeholder="Your name">
  </div>
  <div class="form-group">
    <label for="gender">Gender</label>
    <input type="text" name="gender" class="form-control" id="gender" placeholder="Gender">
  </div>
  <div class="form-group">
    <label for="birthday">Birthday</label>
    <input type="date" name="birthday" class="form-control" id="birthday" placeholder="Birthday">
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
  </form>
{{ end }}
