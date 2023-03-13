default: build

build:
	yarn --cwd frontend build
	go build -ldflags "-s -w" .

clean:
	rm data.db
	rm AreYaCoding
	rm -R frontend/dist
