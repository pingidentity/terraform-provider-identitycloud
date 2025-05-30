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
	TF_ACC=1 go test `go list ./internal/... | grep -v -e promotion` -timeout 10m -v -p 4

testaccpromotion:
	TF_ACC=1 TF_LOG=warn go test `go list ./internal/... | grep -e promotion` -timeout 10m -v -p 4

testaccfolder:
	TF_ACC=1 go test ./internal/resource/${ACC_TEST_FOLDER}... -timeout 10m -v -count=1

devchecknotest: install golangcilint generate tfproviderlint tflint terrafmtlint importfmtlint providerschemalint

devcheck: devchecknotest testacc

golangcilint:
	go tool golangci-lint run --timeout 5m ./internal/...

tfproviderlint: 
	go tool tfproviderlintx \
						-c 1 \
						-AT001.ignored-filename-suffixes=_test.go \
						-AT003=false \
						-XAT001=false \
						-XR004=false \
						-XS002=false ./internal/...

tflint:
	go tool tflint --recursive --disable-rule "terraform_unused_declarations" --disable-rule "terraform_required_providers" --disable-rule "terraform_required_version"

terrafmtlint:
	find ./internal/acctest -type f -name '*_test.go' \
		| sort -u \
		| xargs -I {} go tool terrafmt -f fmt {} -v

importfmtlint:
	go tool impi --local . --scheme stdThirdPartyLocal ./internal/...

providerschemalint:
	python3 ./scripts/schemaLint.py
