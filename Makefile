buildDB:
	cd ./LibraryData && $(MAKE) buildDB

startDB:
	cd ./LibraryData && $(MAKE) startDB

first:
	cd ./LibraryService && $(MAKE) first

server_start:
	cd ./LibraryService && $(MAKE) server_start

client_start:
	cd ./LibraryService && $(MAKE) client_start


#================
stopDB:
	cd ./LibraryData && $(MAKE) stop

tests:
	cd ./LibraryService && $(MAKE) tests


