build:
	docker build -t library_d .

run:
	docker run --name library_d -d \
	-p 3307:3307 \
	-e MYSQL_ROOT_PASSWORD=root \
    --restart unless-stopped \
    -v mysql:/var/lib/mysql \
    mysql:latest

