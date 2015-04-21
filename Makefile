build:
	go build -o ./gopher-spree-api-dev ./interfaces/web

run: build
	forego start
