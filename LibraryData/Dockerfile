# Берем за основу последнюю версию докер образа MySQL
FROM mysql:latest

# Добавляем базу
ENV MYSQL_DATABASE library

#Добавляем скрипты в docker-entrypoint-initdb.d/ они будут выполнены автоматически при запуске контеинера
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/