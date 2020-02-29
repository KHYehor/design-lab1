default: out/example

clean:
	rm -rf out build

test:
	go vet && go test

out/example: implementation/implementation.go cmd/example/main.go
	mkdir -p out build
	git describe | xargs -I{} echo -e "package build \nvar BuildVersion string = \"{}\"" > ./build/build.go
	go build -o out/example ./cmd/example