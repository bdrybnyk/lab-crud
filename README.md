# Guitar CRUD API (Eng)

Professional RESTful API. The project follows clean architecture principles, utilizing Go 1.22+ and PostgreSQL.

## Project Journey & Implementation
This project involved a full-cycle development process on Linux, including:
* **Database Hardening**: Configuring PostgreSQL Ident/MD5 authentication for secure local development.
* **Clean Architecture**: Decoupling business logic from the database layer using interfaces for high testability.
* **Modern Routing**: Using native Go 1.22 `http.ServeMux` features for method-based routing and path parameters.
* **Bonus Tasks**:
    * **Viper**: Dynamic configuration via `.env` files.
    * **Golang-migrate**: Automated SQL schema migrations on startup.
    * **Validation**: Strict payload checks using `validator/v10`.

## Tech Stack
* **Language**: Go 1.22
* **Database**: PostgreSQL
* **Configs**: Viper
* **Migrations**: golang-migrate
* **Testing**: Bruno (Client-side) & Go Testing (Unit-tests)

## API Testing with Bruno
The project includes a full collection for **Bruno**
1. Install Bruno via Flatpak: `flatpak install flathub com.usebruno.Bruno`.
2. Open the `Guitar API` folder located in this repository.
3. The collection includes 6 pre-configured requests:
    * **CREATE**: POST with validation (strings 4-12).
    * **GET ALL**: Retrieve the full list.
    * **GET BY ID**: Dynamic UUID lookup.
    * **PUT**: Full entity update.
    * **PATCH**: Toggle `is_electric` status.
    * **DELETE**: Remove record from DB.

## Setup & Run

1. **Environment**:
Create a `.env` file in the root:
```env
DB_URL=postgres://postgres:password@localhost:5432/postgres?sslmode=disable
SERVER_PORT=:8080
```

2. **Run Server**:
```Bash
go run .
```

3. **Run Unit Tests**:
```bash
go test -v
```

---

# Guitar CRUD API  (Ukr)

Професійна RESTful API. Проєкт побудований за принципами чистої архітектури з використанням Go 1.22+ та PostgreSQL.

## Історія розробки та реалізація
Проєкт пройшов повний цикл розробки на Linux, що включав:
* **Конфігурація БД**: Налаштування автентифікації PostgreSQL (Ident/MD5) для стабільної локальної розробки.
* **Архітектура**: Відокремлення логіки обробника від бази даних через інтерфейси.
* **Сучасний роутинг**: Використання нативних можливостей Go 1.22 (`http.ServeMux`) для маршрутизації.
* **Бонусні завдання**:
    * **Viper**: Динамічні конфіги через `.env` файл.
    * **Golang-migrate**: Автоматичні міграції SQL-схеми при старті сервера.
    * **Валідація**: Жорстка перевірка JSON-даних через `validator/v10`.

## Тестування через Bruno
Проєкт містить повну колекцію запитів для **Bruno**
1. Встановлення: `flatpak install flathub com.usebruno.Bruno`.
2. Відкрийте папку `Guitar API`, що лежить у цьому репозиторії.
3. Колекція містить 6 сценаріїв:
    * **CREATE**: Створення з перевіркою валідації (струни 4-12).
    * **GET ALL**: Отримання всього списку.
    * **GET BY ID**: Пошук за UUID.
    * **PUT**: Повне оновлення об'єкта.
    * **PATCH**: Зміна статусу `is_electric`.
    * **DELETE**: Видалення запису.

## Запуск

1. **Оточення**:
Створіть файл `.env` у корені:
```env
DB_URL=postgres://postgres:password@localhost:5432/postgres?sslmode=disable
SERVER_PORT=:8080
```

2. **Запуск**:
```bash
go run .
```

3. **Тести**:
```bash
go test -v
```