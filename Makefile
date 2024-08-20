.PHONY: build clean

build:
	@make clean
	@mkdir -p build/pichub/web
	@cp -r web/* build/pichub/web
	@cp -r config/ build/pichub
	@GOOS=linux GOARCH=amd64 go build -o build/pichub/pichub
	@GOOS=windows GOARCH=amd64 go build -o build/pichub/pichub.exe
	@cd build && zip -rq pichub-windows-amd64.zip pichub/web/ pichub/config/ pichub/pichub.exe  > /dev/null && cd ..
	@tar -czf build/pichub-linux-amd64.tar.gz -C build/ pichub/web/ pichub/config/ pichub/pichub
	@rm -rf build/pichub
	@unset version

clean:
	@rm -rf build

