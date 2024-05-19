<h1 align="center"> eShop - CDN </h1> <br>
<div>
    <img src="assets/logo.png" width="250" height="200" style="display: block;margin-left: auto;margin-right: auto;>
    <hr>
    <p align="center">
      Microservice provide file(s) storage (S3) with real-time subscriptions
    </p>
</div>
<hr>

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Run Lint and Testing](https://github.com/WildEgor/e-shop-support-bot/actions/workflows/lint.yml/badge.svg)](https://github.com/WildEgor/e-shop-cdn/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/WildEgor/e-shop-cdn/branch/develop/graph/badge.svg)](https://codecov.io/gh/WildEgor/e-shop-cdn/branch/develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/WildEgor/e-shop-cdn)](https://goreportcard.com/report/github.com/WildEgor/e-shop-cdn)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/WildEgor/e-shop-cdn)
[![Publish Docker image](https://github.com/WildEgor/e-shop-cdn/actions/workflows/publish.yml/badge.svg)](https://hub.docker.com/repository/docker/wildegor/e-shop-cdn)

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Quick Start](#quick-start)
- [Contributing](#contributing)

## Introduction

Service allow upload multiple files to shared S3 public bucket

## Features

- [x] Upload multiple files;
- [x] Save files metadata to database;
- [x] Delete file;
- [x] Subscribe for files changes;
- [] Replace file;

## Requirements

- [Git](http://git-scm.com/)
- [Go >= 1.22](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Task](https://taskfile.dev/installation/)
- [Air](https://github.com/cosmtrek/air?tab=readme-ov-file#via-go-install-recommended)
- [MongoDB](https://www.mongodb.com/)

## Quick start

1. Start MongoDB using [docker-compose](https://github.com/WildEgor/e-shop-dot/blob/develop/docker-compose.yaml#L130);
2. Start Minio using [docker-compose](https://github.com/WildEgor/e-shop-dot/blob/develop/docker-compose.yaml#L162);
3. Prepare .env file using example:
```text

```
4. Install tools above:
5. Run service using ```air``` or ```docker-compose```:
```shell
task local-dev
```
or
```shell
task docker-dev
```

## Contributing

Please, use git cz for commit messages!
```shell
git clone https://github.com/WildEgor/e-shop-cdn
cd e-shop-cdn
git checkout -b feature-or-fix-branch
git add .
git cz
git push --set-upstream-to origin/feature-or-fix-branch
```

## License

<p>This project is licensed under the <a href="LICENSE">MIT License</a>.</p>

Made with ❤️ by me
