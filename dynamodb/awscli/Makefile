up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

restart:
	docker-compose restart

dev:
	docker-compose up

awscli:
	awscli:
	docker run --rm -it -v ~/.aws:/root/.aws amazon/aws-cli dynamodb create-table \
    --endpoint-url=http://localhost:4566 \
    --table-name MyTable \
    --attribute-definitions AttributeName=ID,AttributeType=S \
    --key-schema AttributeName=ID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1