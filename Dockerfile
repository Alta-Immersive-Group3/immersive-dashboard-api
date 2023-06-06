FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /api
COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE $PORT

CMD [ "./main" ]