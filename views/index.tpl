{{ template "base.tpl" . }}

{{ define "content" }}
  {{range $object := .posts}}
    <article>
      <h2><a href="{{urlfor "PostController.GetOne" ":id" $object.Id}}">{{$object.Title}}</a></h2>
      <div class="content border-top">
        {{$object.Content}}
      </div>
      <div class="footer border-bottom text-right">
        {{$object.Member.Username}} at {{date $object.PostedAt "Y-m-d h:i:s"}}
      </div>
    </article>
  {{end}}

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
