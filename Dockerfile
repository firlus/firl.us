# FROM golang:latest
# RUN mkdir -p /app
# WORKDIR /app
# ADD . /app
# RUN go build ./cmd/firl.us/firl.us.go
# EXPOSE 80 
# CMD ["./firl.us"]


# build stage frontend
FROM node:lts-alpine as build-frontend-stage
WORKDIR /app
COPY web/package*.json ./
RUN npm install
COPY ./web .
RUN npm run build

# build stage backend
FROM golang:latest as build-backend-stage
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./out/firlus-url ./cmd/firl.us/firl.us.go

# production stage
FROM alpine:latest as production-stage
WORKDIR /opt
COPY docker-entrypoint.sh .
COPY --from=build-frontend-stage /app/public /opt/static
COPY --from=build-backend-stage /app/out /opt
EXPOSE 80
ENTRYPOINT ["./docker-entrypoint.sh"]