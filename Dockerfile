FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /api
COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE $PORT

CMD [ "./main" ]