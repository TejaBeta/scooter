# Scooter 🛴 [![build](https://github.com/TejaBeta/scooter/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/TejaBeta/scooter/actions/workflows/build.yml) [![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](./LICENSE)

A simple micro-service to test infra. This acts as a pilot service to check connectivity.

## Setup

### Export PROJECT_ID

```bash
PROJECT_ID=my_project_id
```

### Build Docker image

```bash
docker build . -t eu.gcr.io/$PROJECT_ID/scooter:master.latest
```

### Push Docker image

```bash
docker push eu.gcr.io/$PROJECT_ID/scooter:master.latest
```
