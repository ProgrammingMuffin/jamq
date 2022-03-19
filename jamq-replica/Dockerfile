FROM golang:1.18.0-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o application

FROM alpine:latest

ENV SECRET=testing123
ENV NONCE_TTL=1000
RUN mkdir /app
COPY --from=0 /app/application /app/
RUN chmod +x /app/application
EXPOSE 5000
CMD ["/app/application"]