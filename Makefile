default: build

clean:
	rm -f ./codescanner

build:
	CGO_ENABLED=1 \
	go build -a -o codescanner *.go

build-windows:
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	GOOS=windows \
	CC=x86_64-w64-mingw32-gcc \
	go build -a -o codescanner.exe *.go

build-linux:
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	GOOS=linux \
	CC=gcc \
	go build -a -o codescanner.linux *.go

test:
	go test ./...
