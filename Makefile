all: build

build:
	templ generate
	go build -o vinylbase cmd/web/*.go
	go build -o vinyl-cli cmd/cli/*.go

linuxbuild:
	templ generate
	GOOS=linux GOARCH=amd64 go build -o vinylbase.linux cmd/web/*.go
	GOOS=linux GOARCH=amd64 go build -o vinyl-cli.linux cmd/cli/*.go

clean:
	rm -rf vinylbase
	rm -rf vinylbase.*
	rm -rf vinyl-cli
	rm -rf vinyl-cli.*

test:
