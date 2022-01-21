BUILDPATH=$(CURDIR)
BINARY=superheroe-golang-api
MONGO_DIR=mongo

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

mongo_start:
	@docker-compose -f /$(BUILDPATH)/${MONGO_DIR}/docker-compose.yaml up -d

mongo_stop:
	@docker-compose -f /$(BUILDPATH)/${MONGO_DIR}/docker-compose.yaml stop

mongo_down:
	@docker-compose -f /$(BUILDPATH)/${MONGO_DIR}/docker-compose.yaml down

mongo_prepare:
	./mongo/mongo_script.sh 