build-arm:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/greenskeeper
build-linux:
	env GOOS=linux GOARCH=amd64  go build -o ./dist/greenskeeper
build:
	go build -o ./dist/greenskeeper
watch:
	air
run:
	make build && ./dist/greenskeeper >> logs/greenskeeper.log