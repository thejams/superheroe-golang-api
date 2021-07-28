test: 
	@echo "Ejecutando tests..."
	@go test ./... -v

coverage:
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

mod:
	@echo "Vendoring..."
	@go mod vendor
