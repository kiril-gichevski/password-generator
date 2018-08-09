docker:
	docker build -t password-generator:dev .
	docker run  -p 8000:8000 -d  password-generator:dev

test:
	go test -v ./... -cover
