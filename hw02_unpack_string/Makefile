.PHONY: check
check:
	golangci-lint run .
	go test -v -count=1 -race -timeout=1m .
	@if [ -f ./test.sh ]; then \
		./test.sh; \
	else \
		echo "test.sh not found, skipping"; \
	fi