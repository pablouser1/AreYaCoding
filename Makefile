default: build

build:
	yarn --cwd frontend build
	go build -ldflags "-s -w" .

clean:
	rm -f data.db
	rm -f AreYaCoding
	rm -f AreYaCoding.zip
	rm -rf frontend/dist

package:
	make clean
	make build
	cp .env.example .env
	zip AreYaCoding.zip -r AreYaCoding frontend/dist .env
