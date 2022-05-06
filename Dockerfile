FROM golang:1.17.2-alpine AS builder

COPY ./ /app
WORKDIR /app
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOARCH amd64
ENV GO111MODULE  on
ENV GOPROXY https://goproxy.cn

RUN go build -a -installsuffix cgo -ldflags '-s -w' -o help_center cmd/server/main.go

FROM --platform=linux/amd64 alpine:3.15
COPY --from=builder /app/help_center /
ADD static /static
RUN ln -s /help_center /usr/bin

CMD ["help_center"]