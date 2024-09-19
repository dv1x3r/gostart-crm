# gostart-crm

## Overview

Powerful and simple e-commerce web application starter project.

- Backend: Go with Echo, templ, sqlx;
- Frontend: tailwind, htmx, alpine, w2ui (admin dashboard);
- Live read-only demo: https://democrm.weasel.dev

## Structure

- cmd/server: main entrypoint
- internal/pkg: bootstrapper package
- internal/app/model: shared dto models
- internal/app/component: html templates
- internal/app/endpoint: router handlers
- internal/app/service: service layer
- internal/app/storage: database layer
- migrations: database goose sql migrations
- web: frontend js bundles, css, static and font files

## Prerequisites

- Go >= 1.23
- Bun >= 1.1

## Build

```sh
cp .env.example .env # prepare .env with defaults
bun install # install node dependencies
bun install:tools # install go tools (to the 'build' folder)
bun run build # build everything
./build/server # start the server
bun watch # start the server in watch mode (live reload)
```

## Screenshots

- Admin panel
  [Admin](docs/screenshot-admin.png)

- Client view
  [Client](docs/screenshot-admin.png)
