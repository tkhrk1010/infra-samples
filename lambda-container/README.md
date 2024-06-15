# Lambda container
docker imageでlambdaを動かす

## docs
https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html

## test
```
$ make build
$ make run
$ make post
curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
"Hello !"%     
```
