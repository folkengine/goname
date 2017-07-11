binary = goname

.DEFAULT_GOAL := install

install: clean test
	$(clean)
	@go fmt

test:
	@echo 'Running Tests'
	@go fmt
	@go test
	@go vet
	@golint -set_exit_status
	@echo 'Tests Completed'

clean:
	@rm -f $(binary)
