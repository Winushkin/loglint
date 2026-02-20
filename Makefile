build_plugin:
	go build -buildmode=plugin -o ./loglint.so plugin/main.go

.PHONY: build_plugin install clean

install: dependencies build_plugin clean

dependencies:
	go work init
	go work use .
	git clone https://github.com/golangci/golangci-lint.git && go work use golangci-lint
	cd golangci-lint && git checkout tags/v1.61.0 && go build -o ../golangci ./cmd/golangci-lint && cd ../

clean:
	rm -f go.work go.work.sum
	rm -rf golangci-lint