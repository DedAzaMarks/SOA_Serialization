FROM golang:1.17.7
RUN mkdir formats-comparison
WORKDIR formats-comparison
COPY *.go .
COPY go.mod .
COPY Test.proto .
RUN mkdir protoTest
COPY ./protoTest/Test.pb.go ./protoTest
RUN go mod tidy
CMD ["go", "test", "-bench=.", "-benchmem"]