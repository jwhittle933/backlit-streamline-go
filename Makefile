dump:
	go run ./cmd/mp4/parser/main.go --file ./examples/sample.mp4 --dump
.PHONY: dump

run:
	go run ./cmd/mp4/parser/main.go --file ./examples/sample.mp4
.PHONY: run
