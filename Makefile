################################################################################

# This Makefile generated by GoMakeGen 0.7.1 using next command:
# gomakegen .

################################################################################

.DEFAULT_GOAL := help
.PHONY = fmt all clean deps help

################################################################################

all: yo ## Build all binaries

yo: ## Build yo binary
	go build yo.go

install: ## Install binaries
	cp yo /usr/bin/yo

uninstall: ## Uninstall binaries
	rm -f /usr/bin/yo

deps: ## Download dependencies
	git config --global http.https://pkg.re.followRedirects true
	go get -d -v pkg.re/essentialkaos/ek.v9
	go get -d -v pkg.re/essentialkaos/go-simpleyaml.v1

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

clean: ## Remove generated files
	rm -f yo

help: ## Show this info
	@echo -e '\nSupported targets:\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-12s\033[0m %s\n", $$1, $$2}'
	@echo -e ''

################################################################################
