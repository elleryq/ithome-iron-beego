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
  <a href="{{urlfor "PostController.GetCreatePostForm"}}" class="btn btn-primary">Create</a>
  </div>
  <p>Total records: {{.object_list_len}}</p>
  <table class="table">
    <thead>
      <th scope="col">#</th>
      <th scope="col">Title</th>
      <th scope="col">Member</th>
      <th scope="col">Posted at</th>
      <th scope="col">Modified at</th>
    </thead>
    <tbody>
    {{range $object := .object_list}}
      <tr>
        <td scope="col">{{$object.Id}}</th>
        <td scope="col">{{$object.Title}}</td>
        <td scope="col">{{$object.Member.Username}}</td>
        <td scope="col">{{date $object.PostedAt "Y-m-d h:i:s"}}</td>
        <td scope="col">{{date $object.ModifiedAt "Y-m-d h:i:s"}}</td>
      </tr>
    {{end}}
  </table>
  {{if gt .paginator.PageNums 1}}
  <nav aria-label="Page navigation">
  <ul class="pagination justify-content-end">
    {{if .paginator.HasPrev}}
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkFirst}}">&lt;&lt;</a></li>
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkPrev}}">&lt;</a></li>
    {{else}}
        <li class="page-item disabled"><a class="page-link">&lt;&lt;</a></li>
        <li class="page-item disabled"><a class="page-link">&lt;</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <li{{if $.paginator.IsActive .}} class="page-item active"{{end}}>
            <a class="page-link" href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </li>
    {{end}}
    {{if .paginator.HasNext}}
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkNext}}">&gt;</a></li>
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkLast}}">&gt;&gt;</a></li>
    {{else}}
        <li class="page-item disabled"><a class="page-link">&gt;</a></li>
        <li class="page-item disabled"><a class="page-link">&gt;&gt;</a></li>
    {{end}}
  </ul>
  </nav>
  {{end}}
{{ end }}

{{ define "scripts" }}
{{ end }}
