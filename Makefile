BUILDPATH=$(CURDIR)
BINARY=superheroe-golang-api

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

build: 
	@echo "Compilando..."
	@go build -mod vendor -ldflags "-s -w" -o $(BUILDPATH)/build/bin/${BINARY} src/main.go
	@echo "Binario generado en build/bin/"${BINARY}
