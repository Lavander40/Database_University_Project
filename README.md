docker-compose --build -d

redis
127.0.0.1:4041/get/<id>
127.0.0.1:4041/set + json in body

postgre
127.0.0.1:4042/student/get
127.0.0.1:4041/student/get/<id>
127.0.0.1:4041/attend/rate

elastic
127.0.0.1:4043/lecture/<phrase>