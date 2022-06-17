FROM golang
RUN git clone --depth=1 https://github.com/protocolbuffers/protobuf
RUN apt-get update && apt-get install autoconf libtool -y
RUN cd ./protobuf && ./autogen.sh && ./configure && make && make install && ldconfig
RUN go install github.com/micro/micro/v3@latest
RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN go install github.com/micro/micro/v3/cmd/protoc-gen-micro@master

