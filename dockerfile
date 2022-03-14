 FROM golang:1.17.8-alpine3.15    

 RUN mkdir /app
 WORKDIR /app

 COPY  go.mod go.sum
 RUN go mod download 

 COPY . .

 LABEL  maintainer="Ara Chobanyan <test@email.com>" \
        version="1.0" 

 CMD go run ./cmd/api
