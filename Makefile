hello:
	echo "hello"

clean:
	rm -rf bin/*

build: clean
	go build -o bin/sample -v .

run: build
	bin/sample

air:
	air
