.PHONY: build clean

build:
	@mkdir -p build/pichub/static
	@cp -r static/* build/pichub/static
	@GOOS=windows GOARCH=amd64 go build -o build/pichub/pichub
	@GOOS=linux GOARCH=amd64 go build -o build/pichub/pichub.exe
	@tar -czf build/pichub-windows-amd64.tar.gz -C build/ pichub/static/ pichub/pichub.exe
	@tar -czf build/pichub-linux-amd64.tar.gz -C build/ pichub/static/ pichub/pichub
	@rm -rf build/pichub

clean:
	@rm -rf build

