########################################################################################

.PHONY = fmt all clean deps

########################################################################################

all: yo

yo:
	go build yo.go

deps:
	go get -v pkg.re/essentialkaos/ek.v7
	go get -v pkg.re/essentialkaos/go-simpleyaml.v1

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;

clean:
	rm -f yo

########################################################################################

