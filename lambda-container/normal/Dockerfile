FROM golang:1.22 as build
WORKDIR /helloworld
# Copy dependencies list
COPY lambda/go.mod lambda/go.sum ./
RUN go mod download
# Build with optional lambda.norpc tag
COPY lambda/ .
RUN go build -tags lambda.norpc -o main cmd/main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /helloworld/main ./main
ENTRYPOINT [ "./main" ]
