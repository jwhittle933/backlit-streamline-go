dump:
	@go run ./cmd/mp4/parser/main.go --file ./examples/sample.mp4 --dump
.PHONY: dump

file ?= ./examples/sample.mp4
run:
	@go run ./cmd/mp4/parser/main.go --file $(file)
.PHONY: run
