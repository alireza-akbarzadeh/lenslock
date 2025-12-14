# Makefile for managing Go project scripts

.PHONY: build run clean air
	@echo "Available make commands:"; \
	echo "  build     - Build the Go project"; \
	echo "  run       - Build and run the project"; \
	echo "  clean     - Remove built binaries"; \
	echo "  dev       - Run with Air live reload (requires Air)"; \
	echo "  kill8080  - Kill process using port 8080"; \
	echo "  help      - Show this help message"
	@echo "Available make commands:"; \
	echo "  build     - Build the Go project"; \
	echo "  run       - Build and run the project"; \
	echo "  clean     - Remove built binaries"; \
	echo "  dev       - Run with Air live reload (requires Air)"; \
	echo "  kill8080  - Kill process using port 8080"; \
	echo "  help      - Show this help message"
help:
	@echo "Available make commands:"; \
	echo "  build     - Build the Go project"; \
	echo "  run       - Build and run the project"; \
	echo "  clean     - Remove built binaries"; \
	echo "  dev       - Run with Air live reload (requires Air)"; \
	echo "  kill8080  - Kill process using port 8080"; \
	echo "  killport  - Kill process using a specified port (make killport PORT=4000)"; \
	echo "  help      - Show this help message"
killport:
	@if [ -z "$(PORT)" ]; then \
		echo "Please specify a port: make killport PORT=4000"; \
		exit 1; \
	fi; \
	PID=$(shell lsof -ti :$(PORT)); \
	if [ "$$PID" != "" ]; then \
		echo "Killing process on port $(PORT) (PID: $$PID)"; \
		kill $$PID; \
	else \
		echo "No process found on port $(PORT)"; \
	fi

build:
	go build -o tmp/main main.go

run: build
	./tmp/main

clean:
	rm -rf tmp/main

# Live reload with Air (requires Air to be installed)
dev:
	air

# Kill process using port 8080
kill8080:
	@PID=$(shell lsof -ti :8080); \
	if [ "$$PID" != "" ]; then \
		echo "Killing process on port 8080 (PID: $$PID)"; \
		kill $$PID; \
	else \
		echo "No process found on port 8080"; \
	fi
