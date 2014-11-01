PROGRAM=$(shell basename `pwd`)
VERSION=$(shell git describe --always --tag)

%: basename=${PROGRAM}_${VERSION}_$(shell echo $@ | tr / _)
%:
	install -d dist
	gox -output build/${basename}/{{.Dir}} -osarch="$@"
	install -m 0644 README.md build/${basename}/README.md
	(cd build; zip -r ../dist/${basename}.zip ${basename})
