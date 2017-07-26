### Multi-stage build
FROM golang:1.8.3-alpine3.6 as build

RUN apk --no-cache add git curl openssh

COPY keys/id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa && \
    echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config && \
    git config --global url."ssh://git@github.com:".insteadOf "https://github.com"

RUN go get -u github.com/goadesign/goa/... && \
    go get -u gopkg.in/mgo.v2
RUN git clone git@github.com:JormungandrK/microservice-tools.git /go/src/microservice-tools && \
    cd /go/src/microservice-tools && \
    git checkout -f task-7

COPY . /go/src/user-microservice
RUN go install user-microservice


### Main
FROM alpine:3.6

COPY --from=build /go/bin/user-microservice /usr/local/bin/user-microservice
COPY config.json /config.json
EXPOSE 8080

ENV SERVICE_CONFIG_FILE="config.json"
ENV API_GATEWAY_URL="http://localhost:8001"

CMD ["/usr/local/bin/user-microservice"]
