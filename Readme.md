#### Go Server Example
Simple go server demonstrating How to accept a body in GET request

#### Usage
1. Run the server. `go run main.go`
2. Send the HTTP Request
```
curl -X GET "http://localhost:3000/api?key=value" \
-H "Content-Type: application/json" \
-d '{"users": [{"name": "", "Email": "email"}, {"name": "", "Email": "email"}] }'
```