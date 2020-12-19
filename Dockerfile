FROM golang:1.15
RUN GO111MODULES=on CGO_ENABLED=0 GOOS=linux go get -a github.com/mije/echo/cmd/echo

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/echo /bin/echo
EXPOSE 8080
ENTRYPOINT ["/bin/echo"]