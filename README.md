# CART API
## How to use 
You have many different ways to run the API
- Terminal
```sh
$ curl -X GET "http://localhost:8080/getItems"

$  curl -X POST "http://localhost:8181/addItem" -d "{\"name\":\"test1\",\"price\":123.00}"
```
- Postman
```Plaintext
localhost:8181/addItem
```
Add payload like 
```Plaintext
{"name":"test1","price":123.00}
```
at the window Body and select raw option 

- With Makefile commands
```sh
$ make run
$ make testPost
$ make testGet
```
- For run all the test suite just run
```sh
$ make run
$ make test
```