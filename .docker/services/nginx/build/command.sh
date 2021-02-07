#!/bin/bash

certDir=".docker/services/nginx/certs"

if [ ! -f "$certDir/ssl.key" ]; then
    mkdir -p $certDir
    openssl genrsa 2048 > "$certDir/ssl.key"
    openssl req -new -x509 -nodes -days 365 -subj "/C=VN/ST=Da Nang/L=Da Nang City/O=ABC/OU=ANC/CN=ABC" -key "$certDir/ssl.key" -out "$certDir/ssl.crt"
    chmod 744 "$certDir/ssl.key"
    chmod 744 "$certDir/ssl.crt"
fi

nginx -g "daemon off;"

