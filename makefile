run:
	go run main.go

testing:
	go test ./...

update:
	git add .
	git commit -m "$(commit)"
	git push -u origin master
	