build:
	docker build --platform linux/amd64 -t docker-image:test ./lambda

run:
	docker run -d -p 9000:8080 \
	--entrypoint /usr/local/bin/aws-lambda-rie \
	docker-image:test ./main

post:
	curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'

post_payload:
	curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":"World"}'

ps:
	docker ps

stop:
	docker ps -q | xargs docker stop
