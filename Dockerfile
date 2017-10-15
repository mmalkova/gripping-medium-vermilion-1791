FROM golang:alpine

WORKDIR /go/src/app

#RUN apk add --no-cache git \
#    && apk del git

COPY . .
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]

EXPOSE 8080
