# LiteFS / PostgreSQL Example Application

This repository is a fork of the original LiteFS example application. It has been adapted to also support running with PostgreSQL, allowing you to easily benchmark and compare both setups. You can test everything locally using Docker Compose.

## Application Overview

This application is a simple web service that demonstrates database operations using either LiteFS or PostgreSQL. It exposes three endpoints, each running on a different port:

- **http://localhost:8080:** NGINX Load Balancer
- **http://localhost:8081:** Primary application instance
- **http://localhost:8082:** Secondary application instance (LiteFS: read-only)

You can perform basic CRUD operations on each page and observe how data is synchronized across instances, depending on whether LiteFS or PostgreSQL is used as the backend. This setup helps you compare replication, failover, and consistency behaviors between the two database solutions.

### Load Balancing

Both setups use a load balancer to distribute incoming requests. With LiteFS, the load balancer also forwards write (POST) requests to the primary instance to ensure data consistency, while read requests can be handled by any instance. In contrast, PostgreSQL manages replication internally, so the load balancer only distributes requests eavenly.

## Installation

You can test the app locally using Docker Compose.

### Prerequisites

- Docker installed on your system
- Docker Compose installed

\*\*Note: commands should be run from the top-level directory

### Running with LiteFS

To start the application using LiteFS, run:

```bash
docker-compose -f docker-compose.litefs.yml up
```

### Running with PostgreSQL

To start the application using a standard PostgreSQL setup, run:

```bash
docker-compose -f docker-compose.postgres.yml up
```
