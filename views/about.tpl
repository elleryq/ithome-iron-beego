<!DOCTYPE html>

<html>
<head>
  <title>About</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <h1>About</h1>
  <p>
  Hello, {{.Name}}
  </p>
  <p>
    {{ .Message }}
  </p>
  <p>
    {{ i18n .Lang "Dog is on the piano" }}
  </p>
</body>
</html>
