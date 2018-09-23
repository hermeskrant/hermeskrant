binary = antibuild
templateargs = --config config.json
templaterefresh = $(templateargs) --development --host

install:
	go get -u -v gitlab.com/antipy/antibuild
	go install -v gitlab.com/antipy/antibuild
	go build -o articlecompiler

build: 
	make clean
	./articlecompiler
	$(binary) $(templateargs)

refresh:
	./articlecompiler
	$(binary) $(templaterefresh)

clean:
	rm -rf public

netlify_build: install semantic-netlify
	make clean
	./articlecompiler
	$(GOPATH)/bin/$(binary) $(templateargs)

semantic: $(shell find semantic/semantic/src -type f)
	cd semantic && \
		make build

semantic-netlify: $(shell find semantic/semantic/src -type f)
	cd semantic && \
		make build-netlify