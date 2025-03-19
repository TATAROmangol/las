FROM golang:1.23.1-alpine3.20

RUN apk add --no-cache \
    git \
    bash \
    unzip \
    wget \
    make

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v30.1/protoc-30.1-linux-aarch_64.zip \
    && unzip protoc-30.1-linux-aarch_64.zip -d /usr/local/ \
    && chmod +x /usr/local/bin/protoc \
    && rm protoc-30.1-linux-aarch_64.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

ENV PATH="${PATH}:${GOPATH}/bin"

WORKDIR /lms

COPY go.mod go.sum ./ 
RUN go mod download

COPY . ./
RUN make build 

EXPOSE 8080
EXPOSE 50051

CMD [ "./bin/main" ]