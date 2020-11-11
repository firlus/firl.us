FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build ./cmd/firl.us/firl.us.go
EXPOSE 80 
CMD ["./firl.us"]