FROM library/golang:1.14.0 as build

COPY cmd/server.go /
RUN CGO_ENABLED=0 GOOS=linux go build -o /server /server.go

FROM scratch as final
COPY --from=build /server /
ENTRYPOINT ["/server"]