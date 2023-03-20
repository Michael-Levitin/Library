# Library
Authors and books

(Makefile в папке LibraryData, или общий - в корне проекта)
1. Собрать docker образ базы данных - "buildDB" 
2. Запустить docker образ базы данных - "startDB"

(Makefile в папке LibraryService, или общий - в корне проекта)
3. При первом запуске проверить/обновить зависимости или запустить - "first" 
4. Запустить server сервиса библиотеки - "server_start"
5. Запустить клиент сервиса библиотеки - "client_start" 
6. Запустить тесты сервиса библиотеки - "test"