<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<title>BuildBox аутентификация</title>
</head>
<body>
	{{if .error}}
      {{.error}}
	{{else}}
		<script type="text/javascript">
			{{if .frontcallbackparams}}
				var callback = '{{.frontcallback}}', params = '{{.frontcallbackparams}}';
				{{if .clearparams}}
					callback = callback.replace(/\b({{.clearparams}})=[^&]*&?/ig, '').replace(/\?$/, '');
            {{end}}
				var redirect = /\?/.test(callback) ? callback+'&'+params : callback+'?'+params;
				try {
					window.opener.location.href = redirect;
					window.close();
					/*setTimeout(function () {
						window.opener.location.href = callback;
						window.close();
					}, 500);*/
				} catch (e) {
					document.write('Ошибка обновления страницы. Закройте это окно и обновите страницу.');
				}
			{{else}}
				window.close();
	      {{end}}
		</script>
	{{end}}
</body>
</html>