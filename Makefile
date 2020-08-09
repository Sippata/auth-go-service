all:
	go build -o bin/auth-go github.com/Sippata/auth-go-service/cmd/app
	bin/auth-go