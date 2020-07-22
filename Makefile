src = $(shell find . -type f -name '*.go') $(shell find . -type d \( -path './pkg*' -o -path './cmd*' \))

huecfg: $(src)
	go build

.PHONY: clean
clean:
	rm -f ./huecfg

.PHONY: clean-gen
clean-gen:
	find . -type f -name '*_gen.go' -delete

.PHONY: gen
gen:
	go generate ./...

.PHONY: test
test:
	go test -v -coverprofile .coverage.out ./...
