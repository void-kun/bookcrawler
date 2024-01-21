FROM nginx:1.15

LABEL maintainer="hoangzrik@gmail.com"

RUN apt-get update && apt-get install -y \
        g++ \
        zip \
        vim \
        curl \
        openssl \
        libssl-dev \
        --no-install-recommends apt-utils \
        && rm -r /var/lib/apt/lists/*

WORKDIR /etc/nginx
