# Default url: http://localhost:7331
templ:
	templ generate --watch --proxy="http://localhost:8000" --cmd="air" --open-browser=false -v
tailwind:
	npx --yes tailwindcss -i ./assets/static/css/input.css -o ./assets/static/css/output.css --minify --watch