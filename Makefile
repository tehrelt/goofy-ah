build:
	go build

run:
	make build
	.\goofy-ah.exe -file test.txt
