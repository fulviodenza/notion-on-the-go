FROM golang:1.19.8 as builder

# Create and change to the app directory.
WORKDIR /app
ADD . /app

RUN go build -o /notion-on-the-go

CMD [ "/notion-on-the-go" ]