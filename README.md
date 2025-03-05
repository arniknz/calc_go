# calc_go: Web-service "Calculator"

### Description
This project implements a web service that evaluates arithmetic expressions submitted by the user via an HTTP request or cURL.

## Starting the service
  1. Check that you have [Go](https://go.dev/dl/) installed
  2. Clone the project from GitHub:
     ```
      git clone https://github.com/arniknz/calc_go.git
     ```
  3. To start service
     ```
      cd .\calc_go\
      go run cmd/orchestrator/main.go
      go run cmt/agent/main.go
     ```
  4. The service will be available at: http://localhost:8080/api/v1/calculate

## Endpoints:

Добавление вычисления арифметического выражения:
<code>
  curl --location 'localhost/api/v1/calculate' \
  --header 'Content-Type: application/json' \
  --data '{
    "expression": <строка с выражение>
  }'
</code>

Status Codes: 201 - expression accepted, 422 - Invalid data, 500 - Something went wrond

Response Body

{
    "id": unique_id
}

Получение списка выражений
<code>curl --location 'localhost/api/v1/expressions'</code>

Response Body
<code>
"expressions": [
        {
            "id": <идентификатор выражения>,
            "status": <статус вычисления выражения>,
            "result": <результат выражения>
        },
        {
            "id": <идентификатор выражения>,
            "status": <статус вычисления выражения>,
            "result": <результат выражения>
        }
    ]
</code>

Status Codes:
    200 - успешно получен список выражений
    500 - что-то пошло не так


Получение выражения по его идентификатору
<code>curl --location 'localhost/api/v1/expressions/:id'</code>

Status Codes:
    200 - успешно получено выражение
    404 - нет такого выражения
    500 - что-то пошло не так


# Example requests
### cURL:
| METHOD | cURL request | Response | Status code |
| ------ | ------------ | -------- | ----------- |
| POST   |<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2+2*2"  }'</code>|<code>{"result": 6.000000}</code>| 200 |
| POST   |<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2 + "  }'</code>|<code>{"error": "Invalid expression"}</code>| 422 |
| GET    |<code>curl --request GET --url 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2 + 1"  }'</code>|<code>{"error": "Only POST method is allowed"}</code>| 405 |
| POST   |<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "bebebe": "2 + 2"  }'</code>|<code>{"error": "Bad request"}</code>| 400 |

### Internal Server Error: 500 status code
If an internal server error occurs, the service will return an error with status code 500
**<code>{"error": "Internal server error"}</code>**

### Json Postman requests:
| METHOD | Json request | Response | Status code |
| ------ | ------------ | -------- | ----------- |
| POST   | <code>{  "expression": "2+2*2"  }</code>|<code>{"result": 6.000000}</code>| 200 |
| POST   | <code>{  "expression": "2 + "  }</code>|<code>{"error": "Invalid expression"}</code>| 422 |
| GET    | <code>{  "expression": "2 + 1"  }</code>|<code>{"error": "Only POST method is allowed"}</code>| 405 |
| POST   | <code>{  "bebebe": "2 + 2"  }</code>|<code>{"error": "Bad request"}</code>| 400 |
