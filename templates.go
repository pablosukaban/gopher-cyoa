package gophercyoa

const defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
		<style>
			body {
				font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, sans-serif;
				background: #f4f6f8;
				color: #2d3436;
				margin: 0;
				padding: 40px 20px;
				line-height: 1.6;
			}

			h1 {
				text-align: center;
				margin-bottom: 30px;
				font-weight: 600;
			}

			.container {
				max-width: 700px;
				margin: 0 auto;
				background: white;
				padding: 40px;
				border-radius: 12px;
				box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
			}

			p {
				margin-bottom: 16px;
			}

			ol {
				padding-left: 20px;
			}

			li {
				margin-bottom: 12px;
			}

			a {
				text-decoration: none;
				color: #0984e3;
				font-weight: 500;
				transition: all 0.2s ease;
			}

			a:hover {
				color: #0652DD;
				text-decoration: underline;
			}

			.footer {
				margin-top: 30px;
				font-size: 14px;
				color: #636e72;
				text-align: center;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>{{.Title}}</h1>
			{{range .Story}}<p>{{ . }}</p>{{end}}
			{{if .Options}}
			<ol>
				{{range .Options}}
				<li>
					<a href="{{.Arc}}">
						{{.Text}}
					</a>
				</li>
				{{end}}
			</ol>
			{{else}}
			<a href="intro">
			Home
			</a>
			{{end}}
		</div>
	</body>
</html>`

const StoryHandlerTmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
		<style>
			body {
				font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, sans-serif;
				background: #f4f6f8;
				color: #2d3436;
				margin: 0;
				padding: 40px 20px;
				line-height: 1.6;
			}

			h1 {
				text-align: center;
				margin-bottom: 30px;
				font-weight: 600;
			}

			.container {
				max-width: 700px;
				margin: 0 auto;
				background: white;
				padding: 40px;
				border-radius: 12px;
				box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
			}

			p {
				margin-bottom: 16px;
			}

			ol {
				padding-left: 20px;
			}

			li {
				margin-bottom: 12px;
			}

			a {
				text-decoration: none;
				color: #0984e3;
				font-weight: 500;
				transition: all 0.2s ease;
			}

			a:hover {
				color: #0652DD;
				text-decoration: underline;
			}

			.footer {
				margin-top: 30px;
				font-size: 14px;
				color: #636e72;
				text-align: center;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>{{.Title}}</h1>
			{{range .Story}}<p>{{ . }}</p>{{end}}
			{{if .Options}}
			<ol>
				{{range .Options}}
				<li>
					<a href="/story/{{.Arc}}">
						{{.Text}}
					</a>
				</li>
				{{end}}
			</ol>
			{{else}}
			<a href="intro">
			Home
			</a>
			{{end}}
		</div>
	</body>
</html>`
