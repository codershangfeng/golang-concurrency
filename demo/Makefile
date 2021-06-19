.PHONY: fmt
fmt:
	go fmt ./...


.PHONY: test
test: fmt
	go test -timeout 30s ./...

.PHONY: run
run: fmt
	go run ./main.go