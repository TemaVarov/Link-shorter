
# Link Shorter

Сервис сокращения ссылок. Проект демонстрирует реализацию полноценного backend-приложения на Go с Docker-инфраструктурой, миграциями, авторизацией и хранением статистики переходов.

## Основной функционал
- Создание коротких ссылок
- Переход по сокращённой ссылке и подсчёт статистики
- Авторизация пользователей
- Хранение данных в PostgreSQL
- Автоматические миграции

## Архитектура проекта
<html>
/cmd            - входная точка backend-сервера<br>
/configs        - конфигурация<br>
/internal<br>
    /auth       - авторизация: HTTP-handlers, логика, модель<br>
    /link       - ссылки: HTTP-handlers, логика, модель, работа с репозиторием<br>
    /stat       - сбор статистики переходов: HTTP-handlers, логика, модель, работа с репозиторием <br>
    /user       - модель пользователя + работа с репозиторием<br>
/migrations     - auto-migrate сервис<br>
/pkg            - подключение к бд, di, шина событий, jwt, middleware, работа с запросом и ответом<br>
</html>

## Используемые технологии
- Go (Golang) — бизнес-логика и API
- Gorm — ORM для работы с PostgreSQL
- PostgreSQL — хранение ссылок, пользователей и статистики
- Docker + docker-compose — контейнеризация
- AutoMigrate с retry — автоматическая миграция и создание таблиц при первом запуске
- E2e и unit тесты с моком БД и HTTP

## Запуск через Docker
1. Создать .env, указать конфиг базы и secret для JWT. <br>
Пример:<br>
DSN="host=postgres user=postgres password=my_pass dbname=link port=5432 sslmode=disable"<br>
SECRET="MEECAQAwEwYHKoZIzj0CAQYIKoZIzj0DAQcEJzAlAgEBBCCbOb3hQUlp+4SpcmMJnCGCI5ZjAn2KOnGQnTrJT/IlQw=="<br>

2. Запустить контейнеры:
docker-compose up --build

