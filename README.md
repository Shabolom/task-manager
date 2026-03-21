# Создание миграции
winget install golang-migrate

migrate create -ext sql -dir build/app/migrations <migration_name_change_me>
# Заполнить SQL-команды для создания и отката миграций в build/app/migrations

# Драйвер для ПГ (нужен для работы с базой данных)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Применение миграции (на локальную БД)
migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path build/app/migrations up
