FROM golang:1.24.2

WORKDIR /usr/cmd/main

#RUN go install github.com/air-verse/air@latest
#	go install go.uber.org/mock/mockgen@latest && \
#	go install gotest.tools/gotestsum@latest

# Add GOPATH/bin to PATH
ENV PATH="/go/bin:${PATH}"

COPY .. .
RUN go mod tidy