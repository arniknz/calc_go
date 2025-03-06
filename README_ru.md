# Веб-сервис "Калькулятор"
работает, но без docker'а ай донт ноу

[Английская версия](README.md) <--

### Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, присланные пользователем через HTTP-запрос.

# Запуск сервиса
  1. Проверьте что у Вас установлен [Go](https://go.dev/dl/)
  2. Склонируйте репу:
     ```
      git clone https://github.com/arniknz/calc_go.git
     ```
  3. Чтобы запустить.. запустите:
     ```
      cd .\calc_go\
      go run cmd/orchestrator/main.go
      go run cmt/agent/main.go
     ```
  4. Сервис будет доступен по адресу: ```http://localhost:8080/api/v1/calculate```

# Эндпоинты:

Отправьте выражение:
```
    curl --location 'localhost/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
      "expression": <выражение>
    }'
```

Кода статусов: 201 - Принято для обработки, 422 - Невалидные данные, 500 - Что-то пошло не так

Тело ответа:

```
{
    "id": unique_id
}
```

Получить список всех выражений, которые вы отправили:
```curl --location 'localhost/api/v1/expressions'```

Тело ответа:
```
"expressions": [
        {
            "id": <id>,
            "status": <статус: pending/progress/completed>,
            "result": <результатt>
        },
        {
            "id": <id>,
            "status": <статус: pending/progress/completed>,
            "result": <результат>
        }
    ]
```

Коды статусов:
```
    200 - OK
    500 - Что-то пошло не так
```

Получить выражение по айдишнику:
```curl --location 'localhost/api/v1/expressions/:id'```

Коды статусов:
```
    200 - OK
    404 - Выражение не найдено
    500 - Что-то пошло не так
```

# Экспорт переменных окружения через консоль
```
  export TIME_ADDITION_MS=50
  export TIME_SUBTRACTION_MS=50
  export TIME_MULTIPLICATIONS_MS=100
  export TIME_DIVISIONS_MS=100
  export COMPUTING_POWER=4

```


# Примеры запросов
### cURL:
| МЕТОД | cURL запрос | Ответ | Код статуса |
| ------ | ------------ | -------- | ----------- |
| POST   |<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2+2*2"  }'</code>|<code>{"result": 6.000000}</code>| 200 |
| POST   |<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2 + "  }'</code>|<code>{"error": "Invalid expression"}</code>| 422 |
| GET    |<code>curl --request GET --url 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2 + 1"  }'</code>|<code>{"error": "Only POST method is allowed"}</code>| 405 |
| POST   |<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "bebebe": "2 + 2"  }'</code>|<code>{"error": "Bad request"}</code>| 400 |

### Internal Server Error: 500 код статуса
Если случилось что-то в сервере и это не проблема клиента, то вернется ошибка 500
**<code>{"error": "Internal server error"}</code>**

### Json Postman запросы:
| МЕТОЖ | Json запрос | Ответ | Код статуса |
| ------ | ------------ | -------- | ----------- |
| POST   | <code>{  "expression": "2+2*2"  }</code>|<code>{"result": 6.000000}</code>| 200 |
| POST   | <code>{  "expression": "2 + "  }</code>|<code>{"error": "Invalid expression"}</code>| 422 |
| GET    | <code>{  "expression": "2 + 1"  }</code>|<code>{"error": "Only POST method is allowed"}</code>| 405 |
| POST   | <code>{  "bebebe": "2 + 2"  }</code>|<code>{"error": "Bad request"}</code>| 400 |

## Тестирование: Достаточно просто перейти в директорию с тестами и запустить go test
```
cd test/
go test
```
