#!/bin/bash

ENDPOINT_URL="http://host.docker.internal:14566"
TABLE_NAME="BarkTable"

cd $(dirname "$0") && pwd

# Delete table
docker-compose exec awscli aws dynamodb delete-table \
  --endpoint-url=$ENDPOINT_URL \
  --table-name $TABLE_NAME

# Create table
# localhostじゃないことに注意。
docker-compose exec awscli aws dynamodb create-table \
  --endpoint-url=$ENDPOINT_URL \
  --table-name $TABLE_NAME \
  --attribute-definitions AttributeName=Username,AttributeType=S \
  --key-schema AttributeName=Username,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

docker-compose exec awscli aws dynamodb update-table \
  --endpoint-url=$ENDPOINT_URL \
  --table-name $TABLE_NAME \
  --stream-specification StreamEnabled=true,StreamViewType=NEW_IMAGE
