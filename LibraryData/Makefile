buildDB:
	docker build -t library .

startDB:
	docker run -d -p 3307:3306 --name library \
	-e MYSQL_ROOT_PASSWORD=password \
	--restart unless-stopped \
    library

stopDB:
	docker stop library
	docker rm library
