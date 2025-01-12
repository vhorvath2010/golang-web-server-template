dev:
	templ generate --watch & air

tailwind:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch 