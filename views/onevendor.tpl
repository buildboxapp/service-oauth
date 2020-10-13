<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>BuildBox аутентификация</title>
</head>
<body>
	<script type="text/javascript">
      {{range .buttons}}
		   document.location.href = '{{.Link}}';
      {{end}}
	</script>
</body>
</html>