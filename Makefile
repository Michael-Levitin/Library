buildDB:
	cd ./LibraryData && make build

runDB:
	cd ./LibraryData && make run

server_start:
	cd ./LibraryService && make server

client_start:
	cd ./LibraryService && make client


#================
stopDB:
	cd ./LibraryData && make stop


