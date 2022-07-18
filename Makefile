all: run

dump:
	@go run ./cmd/mp4/parser/main.go --file ./examples/sample.mp4 --dump
.PHONY: dump

file ?= "$(HOME)/Desktop/DadMusic/01FirstDayInHeaven.m4a"
run:
	@go run ./cmd/mp4/parser/main.go --file $(file)
.PHONY: run
