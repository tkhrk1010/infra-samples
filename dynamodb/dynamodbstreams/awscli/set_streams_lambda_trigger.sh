#!/bin/bash

ENDPOINT_URL="http://host.docker.internal:14566"
FUNCTION_NAME="publishNewBark"

cd $(dirname "$0") && pwd

# ストリームのリストを取得
STREAMS=$(docker-compose exec awscli \
	aws dynamodbstreams list-streams \
		--endpoint-url $ENDPOINT_URL)

# 最初のストリームのARNを取得
STREAM_ARN=$(echo $STREAMS | jq -r '.Streams[0].StreamArn')

echo "Stream ARN: $STREAM_ARN"

docker-compose exec awscli aws lambda create-event-source-mapping \
	--endpoint-url=$ENDPOINT_URL \
  --function-name $FUNCTION_NAME \
  --event-source-arn $STREAM_ARN \
  --starting-position LATEST
