FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /api
COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE $JWT_KEY
EXPOSE $DBUSER
EXPOSE $DBPASS
EXPOSE $DBHOST
EXPOSE $DBPORT
EXPOSE $DBNAME

CMD [ "./main" ]