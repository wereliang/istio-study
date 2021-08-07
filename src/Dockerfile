FROM golang:1.13 AS builder
WORKDIR /home/service
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o basicinfo -tags netgo basicinfo.go   \
    && CGO_ENABLED=0 GOOS=linux go build -o query -tags netgo query.go   \
    && CGO_ENABLED=0 GOOS=linux go build -o recommend_v1 -tags netgo recommend_v1.go \
    && CGO_ENABLED=0 GOOS=linux go build -o recommend_v2 -tags netgo recommend_v2.go \
    && CGO_ENABLED=0 GOOS=linux go build -o star -tags netgo star.go

FROM scratch as basicinfo
WORKDIR /
COPY --from=builder /home/service/basicinfo .
EXPOSE 8080
ENTRYPOINT ["/basicinfo"]

FROM scratch as query
WORKDIR /
COPY --from=builder /home/service/query .
EXPOSE 8080
ENTRYPOINT ["/query"]

FROM scratch as recommend_v1
WORKDIR /
COPY --from=builder /home/service/recommend_v1 .
EXPOSE 8080
ENTRYPOINT ["/recommend_v1"]

FROM scratch as recommend_v2
WORKDIR /
COPY --from=builder /home/service/recommend_v2 .
EXPOSE 8080
ENTRYPOINT ["/recommend_v2"]

FROM scratch as star
WORKDIR /
COPY --from=builder /home/service/star .
EXPOSE 8080
ENTRYPOINT ["/star"]

