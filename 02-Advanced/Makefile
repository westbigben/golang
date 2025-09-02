build:
	@echo "🔨 Building all advanced projects..."
	@for dir in 01_rest_api 02_cli_tool 03_concurrency 04_dockerized_app; do \

		(cd $$dir && go mod tidy && go build -o app); \

	done
	@echo "✅ Build completed!"

run-rest:
	@cd 01_rest_api && go run main.go

run-cli:
	@cd 02_cli_tool && go run main.go

docker-build:
	@cd 04_dockerized_app && docker build -t go-advanced-app .

docker-run:
	@docker run -p 8080:8080 go-advanced-app

clean:
	@echo "🧹 Cleaning builds..."
	@find . -name "app" -type f -delete
