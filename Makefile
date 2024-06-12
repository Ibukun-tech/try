.PHONY:  all build and run
Bin=Bin

start:
	go run  $(run).go
first:
	 @echo "first running ......."
convert:
	 @echo "now converting into binary file........"
	go build -o $(Bin) ./cmd/.

run: convert
	 @echo run file
	go run ./cmd/*.go

up:	
	@echo "Running docker ......"	
	Docker-compose up
down:
	@echo "shutting docker ......"
	Docker-compose down
migrate:
	@echo "Sending default data to database ......"
	go run ./Populate/main.go