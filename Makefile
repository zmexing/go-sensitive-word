.PHONY:build
build: ## 编译项目
	go build -v ./...

.PHONY:test
test: ## 运行测试
	go test -v ./...

.PHONY:lint
lint: ## 执行代码静态分析
	golangci-lint run

.PHONY:bench
bench: ## 运行基准测试
	go test -benchmem -bench .

.PHONY:help
.DEFAULT_GOAL:=help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'