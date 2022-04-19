build:
	go build content_search.go

clear:
	rm -f content_search

install: build
	cp ./content_search /usr/bin/
