templ: 
	templ generate -watch -proxy=http://localhost:8000
tailwind:
	npx tailwindcss -i ./assets/static/css/input.css -o ./assets/static/css/output.css --watch