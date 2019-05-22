FROM debian:8 as builder

RUN apt-get update && \
    apt-get install -y openssl && \
    openssl genrsa -des3 -passout pass:x -out server.pass.key 2048 && \
    openssl rsa -passin pass:x -in server.pass.key -out server.key && \
    rm server.pass.key && \
    openssl req -new -key server.key -out server.csr \
        -subj "/C=US/ST=NY/L=NYC/O=Good/OU=IT Department/CN=selfsigned.com" && \
    openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

FROM nginx:1.15-alpine

COPY --from=builder /server.* /etc/nginx/
COPY default.conf /etc/nginx/conf.d/

EXPOSE 8443

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]
