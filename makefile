binary = antibuild
templateargs = --config config.json
templaterefresh = $(templateargs) --development --host

install:
	go get -u -v gitlab.com/antipy/antibuild
	go install -v gitlab.com/antipy/antibuild

build: 
	make clean
	$(binary) $(templateargs)

refresh:
	$(binary) $(templaterefresh)

clean:
	rm -rf public

netlify_build: install
	make clean
	$(GOPATH)/bin/$(binary) $(templateargs)
	cp -a static/. public/
