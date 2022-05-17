FROM golang:1.16 as builder
WORKDIR /mywork/mockenvoy/
COPY ./*.go ./
ARG LDFLAGS
RUN echo $(LDFLAGS)
RUN go env -w GOPROXY=http://goproxy.cn,direct && \
    go mod init main && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "${LDFLAGS}" -o mockenvoy *.go

FROM centos:centos7
COPY --from=builder /mywork/mockenvoy/mockenvoy .
EXPOSE 8080
ENTRYPOINT [ "/mockenvoy" ]
