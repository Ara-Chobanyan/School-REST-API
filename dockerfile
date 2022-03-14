 FROM golang:1.17.8-alpine3.15    

 RUN mkdir /app
 ADD . /app
 WORKDIR /app

 COPY go.mod ./
 COPY go.sum ./
 RUN go mod download


 RUN go build ./cmd/api

 EXPOSE 4000

 LABEL  maintainer="Ara Chobanyan <test@email.com>" \
        version="1.0" 

 CMD [ "./api" ] 
