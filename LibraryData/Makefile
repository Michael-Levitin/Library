build:
	docker build -t library .

run:
	docker run -d -p 3307:3306 --name library \
	-e MYSQL_ROOT_PASSWORD=password \
	--restart unless-stopped \
    library

stop:
	docker stop library
	docker rm library
