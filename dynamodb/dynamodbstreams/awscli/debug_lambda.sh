#!/bin/bash

ENDPOINT_URL="http://host.docker.internal:14566"
FUNCTION_NAME="publishNewBark"

cd $(dirname "$0") && pwd

docker-compose exec -T awscli aws logs describe-log-groups \
	--endpoint-url=$ENDPOINT_URL

docker-compose exec -T awscli aws logs describe-log-streams \
	--log-group-name "/aws/lambda/$FUNCTION_NAME" \
	--endpoint-url=$ENDPOINT_URL

docker-compose exec -T awscli aws logs filter-log-events \
	--endpoint-url=$ENDPOINT_URL \
	--log-group-name /aws/lambda/$FUNCTION_NAME \
	--filter-pattern '' \
	--query 'sort_by(events, &timestamp)[-20:].[timestamp, message]' \
	--output text
