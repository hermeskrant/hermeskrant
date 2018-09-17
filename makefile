binary = antibuild
templateargs = --config config.json
templaterefresh = $(templateargs) --development --host

$(binary):
	go get -u -v gitlab.com/antipy/antibuild
	go install -v gitlab.com/antipy/antibuild

build: $(binary)
	./articlecompiler
	make clean
	$(binary) $(templateargs)

refresh: $(binary)
	./articlecompiler
	$(binary) $(templaterefresh)

clean:
	rm -rf public

netlify_build: $(binary)
	./articlecompiler
	make clean
	$(GOPATH)/bin/$(binary) $(templateargs)
	cp -a static/. public/
