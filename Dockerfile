From golang:1.3.1-onbuild

ADD . var/www/go/src/github.com/polluxx/bb_serve

RUN go install github.com/polluxx/bb_serve

ENTRYPOINT var/www/go/bin/bb_serve

EXPOSE 8080
