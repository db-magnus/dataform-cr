FROM golang:1.13 as builder
WORKDIR /app
COPY invoke.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

FROM node:12.18.1
WORKDIR /dataform
COPY dataform.json ./
COPY definitions ./
COPY package.json ./
COPY --from=builder /app/server ./
COPY script.sh ./

RUN npm i -g @dataform/cli
RUN dataform install

ENTRYPOINT "./server"

