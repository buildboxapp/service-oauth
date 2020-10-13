<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link href="/auth/login/static/fontawesome/css/fontawesome.min.css" rel="stylesheet">
	<link href="/auth/login/static/fontawesome/css/brands.min.css" rel="stylesheet">
	<link href="/auth/login/static/fontawesome/css/solid.min.css" rel="stylesheet">
	<link href="/auth/login/static/style.css" rel="stylesheet">
	<title>BuildBox аутентификация</title>
</head>
<body>
	<div class="bboauth-wrapper">
		<div class="bboauth-container">
			<a class="logo" href="https://buildbox.app" target="_blank">&nbsp;</a>
          {{range .buttons}}
				 <a class="bboauth-button {{.AClass}} bboauth-button-{{.Vendor}}" href="{{.Link}}" title="{{.Title}}"><i class="{{.IconClass}}"></i></a>
          {{end}}
		</div>
	</div>
</body>
</html>