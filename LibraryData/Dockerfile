FROM mysql:latest
ENV MYSQL_DATABASE library
COPY tables.sql /docker-entrypoint-initdb.d/
COPY demo.sql /docker-entrypoint-initdb.d/

EXPOSE 3307

CMD ["mysqld", "--port=3307", "--datadir=/var/lib/mysql", "--user=mysql"]




