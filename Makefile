make prod:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o production_shell *.go

make run:
	go run *.go
