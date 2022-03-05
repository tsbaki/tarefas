curl -X POST -H "content-type: application/json" http://localhost:8080/todo -d '{"description":"Do this thing later", "project": "Time machine"}'
curl http://localhost:8080/todos
