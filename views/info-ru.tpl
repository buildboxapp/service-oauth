<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link href="/{{.prefix}}static/fontawesome/css/fontawesome.min.css" rel="stylesheet">
	<link href="/{{.prefix}}static/fontawesome/css/brands.min.css" rel="stylesheet">
	<link href="/{{.prefix}}static/fontawesome/css/solid.min.css" rel="stylesheet">
	<link href="/{{.prefix}}static/style.css" rel="stylesheet">
	<link href="/{{.prefix}}static/info.css" rel="stylesheet">
	<script src="https://code.jquery.com/jquery-3.4.1.min.js"></script>
	<script src="/{{.prefix}}static/login.js"></script>
	<title>Как использовать аутентификацию BuildBox</title>
	<script type="text/javascript">
		$(document).ready(function () {
			$('.login').bblogin();
		});
	</script>
</head>
<body>
	<h1>BuildBox oauth2 аутентификация</h1>

	<!------------------------------------>

	<h2>Используя Vanilla JS</h2>
	<xmp>
<script src="{{.host}}/{{.prefix}}static/login.js"></script>

<button onclick="bblogin('http(s)://path/to/callback')">Login</button>
или
<a href="javascript:;" onclick="bblogin('http(s)://path/to/callback')">Login</a>
	</xmp>
	<p>
		Пример:
		<button onclick="bblogin('')">Login button</button>
		<a href="javascript:;" onclick="bblogin('')">Login link</a>
	</p>

	<!------------------------------------>

	<h2>Используя jQuery</h2>
	<xmp>
<script src="https://code.jquery.com/jquery-3.4.1.min.js"></script>
<script src="{{.host}}/{{.prefix}}static/login.js"></script>
<script type="text/javascript">
   $(document).ready(function () {
      $('.login').bblogin('http(s)://path/to/callback');
   });
</script>

<button class="login">Login</button>
или
<a href="#" class="login">Login</a>
	</xmp>
	<p>
		Пример:
		<button class="login">Login button</button>
		<a href="#" class="login">Login link</a>
	</p>

	<!------------------------------------>

	<h2>Светлая тема кнопок</h2>
	<xmp>
bblogin('http(s)://path/to/callback', '', 'light')
	</xmp>
	<p>
		Пример:
		<button onclick="bblogin('', '', 'light')">Login</button>
	</p>

	<!------------------------------------>

	<h2>Пользовательский список вендоров</h2>
	<p>Доступные вендоры: {{.vendors}}</p>
	<xmp>
bblogin('http(s)://path/to/callback', 'vk,yandex')
	</xmp>
	<p>
		Пример:
		<button onclick="bblogin('', 'vk,yandex')">Login</button>
	</p>

	<!------------------------------------>

	<h2>Один вендор</h2>
	<p>Если выбран только один вендор, то сразу откроется аутентификация выбранным вендором.</p>
	<xmp>
bblogin('http(s)://path/to/callback', 'google')
	</xmp>
	<p>
		Пример:
		<button onclick="bblogin('', 'google')">Login</button>
	</p>

	<!------------------------------------>

	<h2>Встраивание кнопок в страницу</h2>
	<p>Получите код кнопок из <code>{{.host}}/{{.prefix}}?embed=1&redirect=...&vendors=google,yandex&theme=light</code>
		(любой из параметров после <code>redirecturl</code> можно опустить), затем вставьте его в код страницы.
		Не забудьте подключить статику:</p>
	<xmp>
<link href="{{.host}}/{{.prefix}}static/fontawesome/css/fontawesome.min.css" rel="stylesheet">
<link href="{{.host}}/{{.prefix}}static/fontawesome/css/brands.min.css" rel="stylesheet">
<link href="{{.host}}/{{.prefix}}static/fontawesome/css/solid.min.css" rel="stylesheet">
<link href="{{.host}}/{{.prefix}}static/style.css" rel="stylesheet">
<script src="{{.host}}/{{.prefix}}static/login.js"></script>
	</xmp>
	<p>
		Пример:
	    {{tohtml .embedbtns}}
	</p>
</body>
</html>
