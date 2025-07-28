# BookShelf API

REST API для управления личной библиотекой книг на Go с использованием:
- Echo framework
- PostgreSQL
- Docker

## Запуск

1. Скопируйте `.env.example` в `.env` и заполните настройки
```bash
cp .env.example .env
```

2. Запустите сервисы:
```bash
docker-compose up --build
```

API будет доступно на `http://localhost:8080`

## Endpoints

- `POST /api/v1/books` - Создать книгу
- `GET /api/v1/books` - Список книг
- `GET /api/v1/books/{id}` - Получить книгу
- `PUT /api/v1/books/{id}` - Обновить книгу
- `DELETE /api/v1/books/{id}` - Удалить книгу