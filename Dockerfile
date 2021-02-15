FROM golang:1.15

# Author
MAINTAINER Project:k8s-feature-test MXuDong <2941884109@qq.com>

# Copy file
WORKDIR /go/src/app

COPY . .

ENV GOPROXY https://goproxy.cn
# the application envs
ENV SERVICE_IP in_docker
ENV SERVICE_NAME in_docker
ENV SERVICE_NAMESPACE in_docker
ENV IS_IN_CLUSERT false
ENV USE_KUBE_FEATURE false

RUN go get -d -v ./...
RUN go build -o app cmd/main.go

CMD ["./app"]