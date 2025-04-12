# BookStore API

Это API для книжного магазина, разработанное с использованием **Go**, **Gin** и **GORM**. Оно позволяет управлять книгами в магазине, включая создание, получение, обновление и удаление книг. База данных используется **PostgreSQL**.

## Структура проекта

Проект состоит из нескольких основных частей:

- **models**: Содержит структуру данных для книги.
- **repository**: Реализует взаимодействие с базой данных.
- **services**: Содержит бизнес-логику.
- **delivery**: Реализует обработчики HTTP-запросов (контроллеры).
- **routers**: Настройка маршрутов для API.

Структура базы данных 
- Таблица: books

- Поле	Тип	Описание
- id	uint	Идентификатор книги (PK)
- title	string	Название книги
- description	string	Описание книги
- price	float64	Цена книги
- stock	int	Количество на складе

Технологии
- Go: Язык программирования для бэкенда.
- Gin: Веб-фреймворк для Go.
- GORM: ORM для работы с PostgreSQL.
- PostgreSQL: Реляционная база данных.
- Docker: Для контейнеризации приложения и базы данных.


Этот API предоставляет функциональность для управления книгами в книжном магазине.

Базовый URL
```go

http://localhost:8080
```
Получить все книги
```go
URL: /books
Метод: GET
```
Создать книгу
```go

- URL: /books
- Метод: POST
- Описание: Создать новую книгу
- {
- "title": "New Book",
- "description": "A detailed description of the new book",
- "price": 15.99,
-  "stock": 20
- }

```




**4. Миграции (Создание таблиц)**

- Все миграционные SQL-файлы находятся в папке:
internal/migration/

- Пример миграционного файла
```go
-- 0002_create_users_table.up.sql
CREATE TABLE users (
id SERIAL PRIMARY KEY,
username TEXT NOT NULL UNIQUE,
password TEXT NOT NULL
);
```




**5. Authentication and Authorization (Middleware)**

- Примеры API-запросов
- Аутентификация
```go
-- POST auth/register
Content-Type: application/json

{
"username": "admin",
"password": "admin123"
}
```

- Логин
```go
-- POST auth/login
Content-Type: application/json

{
"username": "admin",
"password": "admin123"
}
```

- Получить текущего пользователя
```go
-- GET /me
Authorization: Bearer <ваш_jwt_token>
```


