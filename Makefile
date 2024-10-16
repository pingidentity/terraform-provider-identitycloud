SHELL := /bin/bash

.PHONY: install generate fmt vet test testacc golangcilint tfproviderlint tflint terrafmtlint importfmtlint devcheck

default: install

install:
	go mod tidy
	go install .

generate:
	go generate ./...
	go fmt ./...
	go vet ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

testacc:
	TF_ACC=1 go test ./internal/... -timeout 10m -v -count=1

devcheck: install golangcilint generate tfproviderlint tflint terrafmtlint importfmtlint testacc

golangcilint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m ./internal/...

tfproviderlint: 
	go run github.com/bflad/tfproviderlint/cmd/tfproviderlintx \
									-c 1 \
									-AT001.ignored-filename-suffixes=_test.go \
									-AT003=false \
									-XAT001=false \
									-XR004=false \
									-XS002=false ./internal/...

tflint:
	go run github.com/terraform-linters/tflint --recursive --disable-rule "terraform_unused_declarations" --disable-rule "terraform_required_providers" --disable-rule "terraform_required_version"

terrafmtlint:
	find ./internal/acctest -type f -name '*_test.go' \
		| sort -u \
		| xargs -I {} go run github.com/katbyte/terrafmt -f fmt {} -v

importfmtlint:
	go run github.com/pavius/impi/cmd/impi --local . --scheme stdThirdPartyLocal ./internal/...
