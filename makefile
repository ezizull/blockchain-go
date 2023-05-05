run:
ifndef port
	go run main.go
else
	go run main.go -port $(port)
endif

testing:
	go test ./...

update:
	~/go/bin/swag init
	git add .
	git commit -m "$(commit)"
	git push -u origin master
	