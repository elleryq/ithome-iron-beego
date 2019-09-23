{{ template "base.tpl" . }}

{{ define "content" }}
  <h1>User list</h1>
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

  <div>
  <a href="/myuser/create" class="btn btn-primary">Create</a>
  </div>
  <p>Total records: {{.object_list_len}}</p>
  <table class="table">
    <thead>
      <th scope="col">#</th>
      <th scope="col">Name</th>
      <th scope="col">Gender</th>
      <th scope="col">Birthday</th>
    </thead>
    <tbody>
    {{range $object := .object_list}}
      <tr>
        <th scope="row">{{$object.Id}}</th>
        <td>{{$object.Name}}</td>
        <td>{{$object.Gender}}</td>
        <td>{{date $object.Birthday "Y-m-d"}}</td>
      </tr>
    {{end}}
  </table>
{{ end }}

{{ define "scripts" }}
{{ end }}
