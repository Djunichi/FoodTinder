# FoodTinder
## Test Food "Tinder" Project
This repository contains a Docker Compose setup for running a PostgreSQL database,a MongoDB database, MongoDB Viewer and pgAdmin. Follow the instructions below to get started.

## Prerequisites
- Docker installed on your machine
- Docker Compose installed on your machine

## Components

### 1. Application
- **Ports**: 9000
- **Depends On**:
    - PostgreSQL (Service healthy)
    - 
### 2. Postgres
- **Ports**: 5432
- **Healthcheck**:
    - Test: `pg_isready -U testUser -d testDb`
    - Interval: 10s
    - Timeout: 5s
    - Retries: 5

### 3. MongoDB
- **Ports**: 27017

### 4. pgAdmin
[pgAdmin](https://www.pgadmin.org/) is a web-based administration tool for PostgreSQL.

- **Ports**: 5050 (Mapped to 80 in the container)
- **Environment Variables**:
    - `PGADMIN_DEFAULT_EMAIL`: admin@admin.com
    - `PGADMIN_DEFAULT_PASSWORD`: admin
- **Depends On**: PostgreSQL

### 5. MongoDB Viewer
- **Ports** 8081
- **Environment Variables**:
    - `ME_CONFIG_BASICAUTH_USERNAME`: admin
    - `ME_CONFIG_BASICAUTH_PASSWORD`: admin
- **Depends On**: MongoDB
## Getting Started

1. Clone the repository:
    ```sh
    git clone https://github.com/Djunichi/FoodTinder.git
    cd FoodTinder
    ```

2. Ensure you have the necessary files in place:
- `init.sql` for initializing the PostgreSQL database.

3. Run the Docker Compose setup:
    ```sh
    docker-compose up -d
    ```

4. Access the services:
- PostgreSQL: `localhost:5432`
- Swagger UI: `http://localhost:9000/swagger/index.html`
- pgAdmin: `http://localhost:5050`
- MongoDB: `localhost:27017`
- MongoDB Viewer: `http://localhost:8081/`

## Stopping the Setup

To stop the Docker Compose setup, run:
```sh
docker-compose down
```
## Troubleshooting
If you encounter issues with starting the services, check the Docker Compose logs for errors:
```sh
docker-compose logs
```