.PHONY: run

local:
	@echo "==> Iniciando a aplicação..."
	LOG_LEVEL=info \
	DB_CONN_STRING="host=localhost user=admin password=admin dbname=bjj_db port=5432 sslmode=disable" \
	HTTP_PORT=3000 \
	docker-compose up -d && \
	go run cmd/main.go


run:
	@echo "==> Iniciando a aplicação..."
	docker-compose up -d
	@echo "==> waiting app up..."

build:
	@echo "==> Building app"
	docker build -t zap-int:v0 .

docs:
	@echo "==> Updating swagger"
	swag init

install-tools: 
	@echo "==> Installing gotest"
	@go install github.com/rakyll/gotest@latest
	@echo "==> Installing swaggo"
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "==> Installing staticcheck"
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@echo "==> Installing govulncheck"
	@go install golang.org/x/vuln/cmd/govulncheck@latest

go-checks:
	@echo "Rodando validações de segurança no codigo"
	@echo "==> Running staticcheck"
	@staticcheck ./...
	@echo "==> Running govulncheck"
	@govulncheck ./...