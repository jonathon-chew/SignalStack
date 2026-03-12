<p align="center">
  <img src="assets/signalstack-gopher.png" width="700">
</p>

<h1 align="center">SignalStack</h1>
<p align="center">
Distributed infrastructure services written in Go
</p>

SignalStack is an experimental ecosystem of Go services designed to explore modern backend infrastructure, distributed systems, and developer tooling.

The goal of SignalStack is to build a modular platform composed of small, focused services that can be developed independently while working together as a larger system.

Each service implements a specific piece of backend infrastructure such as authentication, caching, messaging, logging, or API routing.

SignalStack is not intended to be a production framework. Instead, it is a learning platform and experimentation environment for building infrastructure-style services in Go.

---

## Architecture

SignalStack follows a service-oriented architecture where each component is responsible for a single capability.

Example high-level architecture:

```bash
Client
   │
   ▼
Gateway
   │
   ├── Auth
   ├── Shortlink
   ├── Upload
   │
   ├── Cache
   ├── Workers
   │
   ├── Logs
   └── Metrics
```

Services communicate primarily over HTTP APIs, with some background tasks handled by worker processes.

---

## Services

| Service | Description |
|-------|-------------|
| auth | Authentication and identity service |
| cache | In-memory caching service |
| workers | Background job processing |
| gateway | API gateway and routing layer |
| logs | Centralised logging service |
| metrics | Metrics collection and monitoring |
| discovery | Service discovery and registration |
| upload | File upload service |
| webhooks | Webhook ingestion and processing |
| shortlink | URL shortening service |
| chat | Websocket chat service |
| proxy | Reverse proxy / load balancer |
| migrate | Database migration tool |
| bench | HTTP benchmarking tool |
| cli | CLI for interacting with SignalStack |

---

## Goals

SignalStack explores patterns commonly used in modern infrastructure systems.

These include:

- service-oriented architecture
- distributed systems design
- background job processing
- observability (logs and metrics)
- API gateways and routing
- caching strategies
- developer tooling

The project prioritises clarity and experimentation over production readiness.

---

## Development Approach

Each service begins as a minimal viable implementation.

Once the core behaviour works, individual components can evolve independently by improving storage layers, adding caching, introducing observability, or improving performance.

This allows experimentation with different infrastructure patterns without needing to redesign the entire system.

---

## Running Locally

Eventually the full SignalStack platform will be runnable locally using Docker Compose.

```bash
docker compose up
```

This will start the core services required to run the platform.

---

## Repository Structure

```bash
signalstack/
  platform/
  auth/
  cache/
  workers/
  gateway/
  logs/
  metrics/
  discovery/
  upload/
  webhooks/
  shortlink/
  chat/
  proxy/
  migrate/
  bench/
  cli/
```

Each service lives in its own repository and can be developed independently.

---

## Why SignalStack Exists

SignalStack is primarily a learning and exploration project.

It provides a practical environment for experimenting with:

- Go backend services
- distributed systems patterns
- infrastructure tooling
- service architecture

By building small focused services, the project aims to mirror the structure of real-world infrastructure ecosystems.

---

## License

MIT
