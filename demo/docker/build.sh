GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o main main.go &&
docker image build -t demo_app ./docker