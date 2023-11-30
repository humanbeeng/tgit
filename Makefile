build:
	go build -o bin/tgit .

clean:
	rm -rf .tgit/

disp:
	ls -lR .tgit

