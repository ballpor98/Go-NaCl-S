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
