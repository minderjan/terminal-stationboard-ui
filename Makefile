build:
	go build -o  bin/stationboard cmd/stationboard/main.go

run:
	go run main.go

clean:
	rm -r bin

install:
	go install ./cmd/stationboard

compile:
    # 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/stationboard-freebsd-386 cmd/stationboard/main.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/stationboard-darwin-386 cmd/stationboard/main.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/stationboard-linux-386 cmd/stationboard/main.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/stationboard-windows-386.exe cmd/stationboard/main.go
        # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/stationboard-freebsd-amd64 cmd/stationboard/main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/stationboard-darwin-amd64 cmd/stationboard/main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/stationboard-linux-amd64 cmd/stationboard/main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/stationboard-windows-amd64.exe cmd/stationboard/main.go