FROM golang:1.22 as builder
WORKDIR /src
COPY main.go /src
RUN CGO_ENABLED=0 GOOS=linux go build -o hello-app main.go

FROM scratch
COPY --from=builder /src/hello-app /hello-app
EXPOSE 8001
CMD ["/hello-app"]