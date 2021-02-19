FROM golang:1.15

# Author
MAINTAINER Project:k8s-feature-test MXuDong <2941884109@qq.com>

# Copy file
WORKDIR /go/src/app

COPY . .

ENV GOPROXY https://goproxy.cn
# the application envs
ENV KTS_ENV_SERVICE_IP in_docker
ENV KTS_ENV_SERVICE_NAME in_docker
ENV KTS_ENV_SERVICE_NAMESPACE in_docker
ENV KTS_IS_IN_CLUSTER false
ENV KTS_USE_KUBE_FEATURE false

RUN go get -d -v ./...
RUN go build -o app k8s-test-backend/cmd

CMD ["./app"]