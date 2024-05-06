#!/bin/bash

ENDPOINT_URL="http://host.docker.internal:14566"

cd $(dirname "$0") && pwd

docker-compose exec awscli aws dynamodb put-item \
  --endpoint-url=$ENDPOINT_URL \
	--table-name BarkTable \
	--item Username={S="Jane Doe"},Timestamp={S="2016-11-18:14:32:17"},Message={S="Testing...1...2...3"}
