.PHONY: build install clean

build:
	go build -o bin/cmus-waybar-lyrics ./cmd/cmus-waybar-lyrics

install: build
	cp bin/cmus-waybar-lyrics /usr/local/bin/

clean:
	rm -rf bin/

run:
	go run ./cmd/cmus-waybar-lyrics/main.go

dev:
	go run ./cmd/cmus-waybar-lyrics/main.go
