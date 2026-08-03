.PHONY: check
check:
	golangci-lint run .
	go test -v -count=1 -race -timeout=1m .
ifneq ($(wildcard ./test.sh),)
	./test.sh
else
	@echo test.sh not found, skipping
endif
