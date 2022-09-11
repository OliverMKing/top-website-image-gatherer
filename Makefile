.PHONY: build-all
build-all: build-windows-amd64 build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64

.PHONY: build-windows-amd64
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/OliverMKing/twig/cmd.VERSION=${VERSION}" -v -o ./bin/twig-windows-amd64.exe

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/OliverMKing/twig/cmd.VERSION=${VERSION}" -v -o ./bin/twig-linux-amd64

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-X github.com/OliverMKing/twig/cmd.VERSION=${VERSION}" -v -o ./bin/twig-linux-arm64

.PHONY: build-darwin-amd64
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/OliverMKing/twig/cmd.VERSION=${VERSION}" -v -o ./bin/twig-darwin-amd64

.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X github.com/OliverMKing/twig/cmd.VERSION=${VERSION}" -v -o ./bin/twig-darwin-arm64
