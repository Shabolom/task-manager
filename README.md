# Локальный запуск проекта

## Требования

Перед запуском убедитесь, что установлены:

-   Docker
-   Docker Compose
-   Go (рекомендуется версия 1.21+)
-   golang-migrate (для применения миграций)

------------------------------------------------------------------------

## Шаг 1. Поднять PostgreSQL через Docker

Из корня проекта выполните:

``` bash
docker compose up -d
```

Проверить статус контейнеров:

``` bash
docker compose ps
```

Остановить контейнеры:

``` bash
docker compose down
```

------------------------------------------------------------------------

## Шаг 2. Создать файл .env

Создайте файл `.env` в корне проекта.

Пример минимальной конфигурации:

``` env
TASK_MANAGER_POSTGRES_HOST=localhost
TASK_MANAGER_POSTGRES_PORT=5432
TASK_MANAGER_POSTGRES_SERVICE_USERNAME=postgres
TASK_MANAGER_POSTGRES_SERVICE_PASSWORD=postgres
TASK_MANAGER_POSTGRES_SERVICE_DATABASE=postgres
TASK_MANAGER_POSTGRES_PARAMS=sslmode=disable
TASK_MANAGER_POSTGRES_MAX_CONNECTION=10
TASK_MANAGER_POSTGRES_MIN_CONNECTION=0
```

Важно: значения должны совпадать с настройками в docker-compose.yml.

Если проект использует дополнительные обязательные переменные окружения
--- добавьте их в этот файл.

------------------------------------------------------------------------

## Шаг 3. Применить миграции

Если база данных пустая, необходимо выполнить миграции.

Установка golang-migrate:

MacOS:

``` bash
brew install golang-migrate
```

Linux:

``` bash
sudo apt install migrate
```

Применение миграций:

``` bash
migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path build/app/migrations up
```

------------------------------------------------------------------------

## Шаг 4. Запустить приложение

После запуска PostgreSQL и настройки `.env` выполните:

``` bash
go run cmd/main.go
```

------------------------------------------------------------------------

## Проверка запуска

Если всё настроено корректно, приложение будет доступно:

    http://localhost:8080

------------------------------------------------------------------------

## Быстрый запуск с нуля

Полный сценарий запуска:

``` bash
docker compose up -d
migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path build/app/migrations up
go run cmd/main.go
```

------------------------------------------------------------------------

## Остановка приложения

Остановить сервер:

    Ctrl + C

Остановить PostgreSQL:

``` bash
docker compose down
```

------------------------------------------------------------------------

## Возможные проблемы

Если приложение не запускается:

Проверьте:

-   контейнер PostgreSQL запущен
-   переменные окружения совпадают с docker-compose.yml
-   миграции применились без ошибок
-   порт приложения свободен
-   DATABASE_URL корректный
