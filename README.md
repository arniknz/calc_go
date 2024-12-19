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
