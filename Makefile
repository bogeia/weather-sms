.PHONY: install
install:
	go mod tidy
	go mod vendor

.PHONY: build
build: install
	@echo ">>>>>>>>make build<<<<<<<<"
	@go build -ldflags=${flags} -o sms