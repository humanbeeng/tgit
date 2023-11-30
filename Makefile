build:
	go build -o bin/tgit .

clean:
	rm -rf .tgit/

display:
	ls -a -r .tgit
