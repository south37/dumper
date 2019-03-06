NAME    := dumper
LDFLAGS := -ldflags="-s -w"

build:
	go build $(LDFLAGS) -o $(NAME) *.go
