FROM golang

WORKDIR /go/src/github.com/rossus/simple_blockchain
COPY . /go/src/github.com/rossus/simple_blockchain

# RUN go get github.com/davecgh/go-spew/spew && go get github.com/gorilla/mux && go get github.com/joho/godotenv