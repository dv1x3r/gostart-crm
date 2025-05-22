# Demo CRM with Go, w2ui and templ

A powerful and simple e-commerce web application starter project.

- Go, Echo, templ, tailwind, htmx, alpine, w2ui (admin dashboard);

## Features

- **Admin dashboard**: Manage the platform through a rich user interface;
- **Client view**: Browse and filter products as a customer;
- **Product management**: Add, update, and manage product listings;
- **Orders management**: Track and process customer orders;

## Build

Make sure you have the following installed to run the project:

- **Go >= 1.24**;

Copy the environment configuration file:

```sh
cp .env.example .env
```

Build and run the project.

Starts the server on `http://localhost:1323`:

```sh
go build -tags \"fts5\" -o ./build/server ./cmd/server/main.go
./build/server
```

Watch mode, which includes air, templ, js and css tracking.

Starts the server with live-reload on `http://localhost:7331`:

```sh
bun watch
```

## Live Demo

- **Live read-only demo**: https://democrm.sbox.weasel.dev

## License

Licensed under the MIT license.
