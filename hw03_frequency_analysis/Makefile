.PHONY: check
check:
	golangci-lint run .
	go test -v -count=1 -race -timeout=1m .
ifeq ($(wildcard ./test.sh),)
	@echo test.sh not found, skipping
else
	bash ./test.sh
endif
