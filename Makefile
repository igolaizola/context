PACKAGE = github.com/igolaizola/context

test:
	go test -v $(PACKAGE) -race

sanity-check:
	golangci-lint run ./...
