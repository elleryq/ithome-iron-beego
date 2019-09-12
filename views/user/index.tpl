{{ template "base.tpl" . }}

{{ define "content" }}
  <h1>User list</h1>
  {{if eq true .has_error}}
  <div class="alert alert-danger" role="alert">
    Error: {{.error}}
  </div>
  {{end}}

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
