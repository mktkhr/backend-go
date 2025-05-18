APP_NAME=server
MAIN := cmd/main.go
OUTPUT := bin/$(APP_NAME)

.PHONY: all build run clean lint help

## build        アプリのバイナリをビルド
build:
	@echo "🔨 Building..."
	@mkdir -p bin
	go build -o $(OUTPUT) $(MAIN)
	@echo "✅ Built at $(OUTPUT)"

## run        アプリのバイナリをビルドして実行
run: build
	@echo "🚀 Running..."
	@./$(OUTPUT)

## lint        Lintの実行(go vet)
lint:
	@echo "🧐 Running go vet..."
	go vet ./...
	@if command -v staticcheck > /dev/null; then \
		echo "🔍 Running staticcheck..."; \
		staticcheck ./...; \
	else \
		echo "⚠️  staticcheck not found."; \
	fi

## clean        ビルド成果物の削除
clean:
	@echo "🧹 Cleaning up..."
	rm -rf bin
	@echo "✅ Cleaned."

## help        ヘルプを表示
help:
	@echo ""
	@echo "🛠️  Available targets:"
	@grep -E '^##' Makefile | awk 'BEGIN {FS = "## "}; {split($$2, a, "        "); printf "  \033[36m%-15s\033[0m %s\n", a[1], a[2]}'
	@echo ""