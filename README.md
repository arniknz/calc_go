# calc_go: Web-service "Calculator"

### Description
This project implements a web service that evaluates arithmetic expressions submitted by the user via an HTTP request or cURL.

## Starting the service
  1. Check that you have [Go](https://go.dev/dl/) installed
  2. Clone the project from GitHub:
     ```
      git clone https://github.com/arniknz/calc_go.git
     ```
  3. Go to your project folder and start the server:
     ```
      cd .\calc_go\
      go run .\cmd\main.go
     ```
  4. The service will be available at: http://localhost:8080/api/v1/calculate


# Example requests
### cURL:
| cURL request | Response | Status code |
| ------------ | -------- | ----------- |
|<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2+2*2"  }'</code>|<code>{"result": 6.000000}</code>| 200 |
|<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2 + "  }'</code>|<code>{"error": "Invalid expression"}</code>| 422 |
|<code>curl --request GET --url 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "expression": "2 + 1"  }'</code>|<code>{"error": "Only POST method is allowed"}</code>| 405 |
|<code>curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{  "bebebe": "2 + 2"  }'</code>|<code>{"error": "Bad request"}</code>| 400 |

### Internal Server Error: 500 status code
If an internal server error occurs, the service will return an error with status code 500

### Json Postman requests:
| Json request | Response | Status code |
| ------------ | -------- | ----------- |
| METHOD = "POST"<code>{  "expression": "2+2*2"  }</code>|<code>{"result": 6.000000}</code>| 200 |
| METHOD = "POST"<code>{  "expression": "2 + "  }</code>|<code>{"error": "Invalid expression"}</code>| 422 |
| METHOD = "GET"<code>{  "expression": "2 + 1"  }</code>|<code>{"error": "Only POST method is allowed"}</code>| 405 |
| METHOD = "POST"<code>{  "bebebe": "2 + 2"  }</code>|<code>{"error": "Bad request"}</code>| 400 |
