# Go-NaCl-S

Golang encryption/decryption(AES-256-GCM) service

## Overview

* [Getting Started](./README.md#getting-started)

## Getting Started

Run with docker compose

* Install Docker
* Install Docker-compose
* Run following command to start service

```bash
docker-compose up --build
```

* call /healthz

```bash
curl -i -X GET 'http://localhost:8080/healthz'
```

* call /encrypt

```bash
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
    '{
    "plaintext": "hello world",
    "nonce": "U2hvTkVXVVNvVVc3"
    }' \
    'http://localhost:8080/encrypt'
```

* call /decrypt

```bash
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
    '{
    "ciphertext": "moY6V5KoQQxoO1tVHt2WGG4MDi4d455W+uxL",
    "nonce": "U2hvTkVXVVNvVVc3"
    }' \
    'http://localhost:8080/decrypt'
```
