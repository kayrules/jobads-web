FROM golang:latest
WORKDIR /Users/kayrules/Projects/go/gopath/src/github.com/kayrules/jobads-web

RUN go get -d -v github.com/facebookgo/grace/gracehttp && \
    go get -d -v github.com/labstack/echo && \
    go get -d -v github.com/gorilla/sessions && \
    go get -d -v github.com/labstack/echo-contrib/session && \
    go get -d -v github.com/globalsign/mgo/bson

ADD ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN ip -4 route list match 0/0 | awk '{print $3 " host.docker.internal"}' >> /etc/hosts
WORKDIR /root/

COPY --from=0 /Users/kayrules/Projects/go/gopath/src/github.com/kayrules/jobads-web/main .
COPY --from=0 /Users/kayrules/Projects/go/gopath/src/github.com/kayrules/jobads-web/view ./view/ 
COPY --from=0 /Users/kayrules/Projects/go/gopath/src/github.com/kayrules/jobads-web/assets ./assets/

EXPOSE 9020
CMD ["./main"]