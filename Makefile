build:
	go build -o bin/tgit .;
	cp bin/tgit ../tgit-test

clean:
	rm -rf .tgit/

disp:
	ls -lR .tgit

