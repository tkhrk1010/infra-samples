#!/bin/bash

ENDPOINT_URL="http://host.docker.internal:14566"

cd $(dirname "$0") && pwd

docker-compose exec awscli aws lambda invoke \
	--cli-binary-format raw-in-base64-out \
	--function-name publishNewBark \
	--payload file:///awscli/payload.json \
	--endpoint-url=$ENDPOINT_URL \
	/awscli/output.txt
