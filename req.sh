curl -X POST http://localhost:8080/addTodo \
-H 'Content-Type: application/json' \
-d '{"project":"ProjectA","Description":"This is an example TODO "}'

curl -X POST http://localhost:8080/addTodo \
-H 'Content-Type: application/json' \
-d '{"project":"ProjectA","Description":"This is another example TODO "}'

curl -X POST http://localhost:8080/addTodo \
-H 'Content-Type: application/json' \
-d '{"project":"ProjectB","Description":"This is an example TODO with another project"}'

curl -X POST http://localhost:8080/addTodo \
-H 'Content-Type: application/json' \
-d '{"project":"ProjectB","Description":"Another TODO"}'
i

