Link Shorter - сервис сокращения ссылок

Проект демонстрирует реализацию полноценного backend-приложения на Go с Docker-инфраструктурой, миграциями, авторизацией и хранением статистики переходов.


Основной функционал

* Создание коротких ссылок
* Переход по сокращённой ссылке и подсчёт статистики
* Авторизация пользователей
* Хранение данных в PostgreSQL
* Автоматические миграции


Архитектура проекта

/cmd            - входная точка backend-сервера
/configs        - конфигурация
/internal
    /auth       - авторизация: HTTP-handlers, логика, модель
    /link       - ссылки: HTTP-handlers, логика, модель, работа с репозиторием
    /stat       - сбор статистики переходов: HTTP-handlers, логика, модель, работа с репозиторием    
    /user       - модель пользователя + работа с репозиторием
/migrations     - auto-migrate сервис
/pkg            - подключение к бд, di, шина событий, jwt, middleware, работа с запросом и ответом


Используемые технологии

* Go (Golang) - бизнес-логика и API
* Gorm - ORM для работы с PostgreSQL
* PostgreSQL - хранение ссылок, пользователей и статистики
* Docker + docker-compose — контейнеризация
* AutoMigrate с retry — автоматическая миграция и создание таблиц при первом запуске
* E2e и unit тесты с моком бд и http


Запуск через Docker

1) Создать .env, указать конфиг базы и secret для jwt
Пример:
DSN="host=postgres user=postgres password=my_pass dbname=link port=5432 sslmode=disable"
SECRET="MEECAQAwEwYHKoZIzj0CAQYIKoZIzj0DAQcEJzAlAgEBBCCbOb3hQUlp+4SpcmMJnCGCI5ZjAn2KOnGQnTrJT/IlQw=="

2) docker-compose up --build

