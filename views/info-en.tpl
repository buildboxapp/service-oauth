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
	<title>How to use BuildBox authentication</title>
	<script type="text/javascript">
		$(document).ready(function () {
			$('.login').bblogin();
		});
	</script>
</head>
<body>
	<h1>BuildBox oauth2 authentication</h1>

	<!------------------------------------>

	<h2>Using Vanilla JS</h2>
	<xmp>
<script src="{{.host}}/{{.prefix}}static/login.js"></script>

<button onclick="bblogin('http(s)://path/to/callback')">Login</button>
or
<a href="javascript:;" onclick="bblogin('http(s)://path/to/callback')">Login</a>
	</xmp>
	<p>
		Example:
		<button onclick="bblogin('')">Login button</button>
		<a href="javascript:;" onclick="bblogin('')">Login link</a>
	</p>

	<!------------------------------------>

	<h2>Using jQuery</h2>
	<xmp>
<script src="https://code.jquery.com/jquery-3.4.1.min.js"></script>
<script src="{{.host}}/{{.prefix}}static/login.js"></script>
<script type="text/javascript">
   $(document).ready(function () {
      $('.login').bblogin('http(s)://path/to/callback');
   });
</script>

<button class="login">Login</button>
or
<a href="#" class="login">Login</a>
	</xmp>
	<p>
		Example:
		<button class="login">Login button</button>
		<a href="#" class="login">Login link</a>
	</p>

	<!------------------------------------>

	<h2>Light buttons theme</h2>
	<xmp>
bblogin('http(s)://path/to/callback', '', 'light')
	</xmp>
	<p>
		Example:
		<button onclick="bblogin('', '', 'light')">Login</button>
	</p>

	<!------------------------------------>

	<h2>Custom vendors list</h2>
	<p>Available vendors: {{.vendors}}</p>
	<xmp>
bblogin('http(s)://path/to/callback', 'vk,yandex')
	</xmp>
	<p>
		Example:
		<button onclick="bblogin('', 'vk,yandex')">Login</button>
	</p>

	<!------------------------------------>

	<h2>One vendor</h2>
	<p>If selected only one vendor, then vendor authentication will directly opened.</p>
	<xmp>
bblogin('http(s)://path/to/callback', 'google')
	</xmp>
	<p>
		Example:
		<button onclick="bblogin('', 'google')">Login</button>
	</p>

	<!------------------------------------>

	<h2>Page embedded buttons</h2>
	<p>Get buttons from <code>{{.host}}/{{.prefix}}?embed=1&redirect=...&vendors=google,yandex&theme=light</code>
		(any params after <code>redirecturl</code> can be omitted), then insert it into page code.
		Don’t forget include static:</p>
	<xmp>
<link href="{{.host}}/{{.prefix}}static/fontawesome/css/fontawesome.min.css" rel="stylesheet">
<link href="{{.host}}/{{.prefix}}static/fontawesome/css/brands.min.css" rel="stylesheet">
<link href="{{.host}}/{{.prefix}}static/fontawesome/css/solid.min.css" rel="stylesheet">
<link href="{{.host}}/{{.prefix}}static/style.css" rel="stylesheet">
<script src="{{.host}}/{{.prefix}}static/login.js"></script>
	</xmp>
	<p>
		Example:
		{{tohtml .embedbtns}}
	</p>
</body>
</html>
