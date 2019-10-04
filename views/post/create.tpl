{{ template "base.tpl" . }}

{{ define "extra_head" }}
<style>
.ck-editor__editable_inline {
    min-height: 400px;
}
</style>
{{end}}

{{ define "content" }}
  <h1>Add new post</h1>
  {{if .flash.error}}
  <div class="alert alert-danger" role="alert">
    Error: {{.flash.error}}
  </div>
  {{end}}
  <form action="" method="POST">
  {{.xsrfdata}}
  <div class="form-group">
    <label for="title">Title</label>
    <input type="text" name="title" class="form-control" id="title" placeholder="Title">
  </div>
  <div class="form-group">
    <label for="content">Content</label>
    <textarea id="content" name="content"></textarea>
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
  </form>
{{ end }}

{{ define "scripts" }}
<script src="https://cdn.ckeditor.com/ckeditor5/12.4.0/classic/ckeditor.js"></script>
<script>
  ClassicEditor
    .create(document.querySelector("#content"), {
		removePlugins: [
		  'EasyImage',
		  'ImageToolbar',
		  'ImageCaption',
		  'ImageStyle',
		  'ImageUpload',
		  'MediaEmbed',
		  'PasteFromOffice'
		],
		image: {}
	})
	.then(editor => {
	  console.log(editor);
	});
</script>
{{ end }}
