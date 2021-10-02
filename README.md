# httpecho-go
An Go/Echo based http mirror. 
Currently is operates on port 1323, custom ports are on the to do list

example request:

`curl http://localhost:1323/get`
`curl -X POST http://localhost:1323/post --header "Content-Type: application/json" --data '{"name":"helloworld"}'`

