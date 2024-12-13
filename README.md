# How is Weather

## Overview

**How is Weather** is a simple Go application that retrieves weather information for a given city or location. It acts as a wrapper around the **Visual Crossing Weather API** and implements a caching strategy using Redis to enhance performance.

This project was developed for learning purposes, specifically to understand how to:

- Integrate Redis with a Go server.
- Implement caching mechanisms to improve response times.
- Containerize processes with Docker and Docker Compose.

## Features

- Retrieve current weather information for a specific city or location.
- Cache weather data in Redis to speed up subsequent requests.
- Automatically fetch fresh weather data from the third-party provider (Visual Crossing) when the cache is empty.
- Simple and clean project structure for easy learning and maintenance.

## Table of Contents

- [How is Weather](#how-is-weather)
  - [Overview](#overview)
  - [Features](#features)
  - [Table of Contents](#table-of-contents)
  - [Pre-requisites](#pre-requisites)
  - [Configuration](#configuration)
    - [Environment Variables](#environment-variables)
  - [Running the Application](#running-the-application)
    - [1. Start Redis using Docker Compose](#1-start-redis-using-docker-compose)
    - [2. Run the Go Application](#2-run-the-go-application)
  - [API Endpoint](#api-endpoint)
    - [GET `/api/weather/{location}`](#get-apiweatherlocation)
      - [Request Example](#request-example)
      - [Response Example](#response-example)
  - [How it Works](#how-it-works)
  - [Docker and Redis Integration](#docker-and-redis-integration)
    - [Docker Compose File](#docker-compose-file)
    - [Key Points:](#key-points)
  - [Visual Crossing Weather API](#visual-crossing-weather-api)
  - [Future Scope](#future-scope)
  - [Summary](#summary)

---

## Pre-requisites

Before running the application, ensure you have the following installed on your system:

1. **Go** (1.19 or higher)

   - Download and install from [Go official website](https://go.dev/).

2. **Docker** and **Docker Compose**

   - Docker is required to run the Redis container.
   - Install Docker from [Docker's official website](https://www.docker.com/).

3. **Visual Crossing Weather API Key**
   - Sign up at [Visual Crossing](https://www.visualcrossing.com/) and obtain an API key.

---

## Configuration

### Environment Variables

- Create a `.env` file in the root directory and populate it like the `.env.example` file.
- Replace `<Your_VisualCrossing_API_Key>` with the API key obtained from Visual Crossing.
- Create a `redis.conf` file inside the root directory and populate ir like the `redis.conf.example` file.

---

## Running the Application

Follow these steps to run the application:

### 1. Start Redis using Docker Compose

Navigate to the root directory of the project and run:

```bash
docker compose up -d
```

This command will:

- Pull the Redis Docker image.
- Start the Redis server in a container.
- Map the Redis port you specified in the `.env` and `redis.conf` to your local machine.

To verify that Redis is running, you can check the logs:

```bash
docker compose logs redis
```

### 2. Run the Go Application

With Redis running, start the Go application:

```bash
go run cmd/how-is-weather/main.go
```

The server will start and listen for incoming API requests on the port you specified inside `.env` file.

---

## API Endpoint

The application exposes a single endpoint to fetch weather data:

### GET `/api/weather/{location}`

**Description**: Retrieve weather information for a specific location.

#### Request Example

```
GET http://localhost:8080/api/weather/london
```

#### Response Example

```json
{
  "address": "london",
  "timezone": "Europe/London",
  "description": "",
  "currentConditions": {
    "datetimeEpoch": 1734106680,
    "temp": 4.1,
    "feelslike": 4.1,
    "humidity": 89.9,
    "conditions": "Overcast"
  }
}
```

---

## How it Works

1. **Client Request**: When the client requests weather data for a specific location:
   - The server first checks Redis for cached data.
2. **Cache Check**:
   - If data is found in Redis, it is returned immediately.
   - If no data is found, the server fetches the data from the **Visual Crossing API**.
3. **Third-party API**: The application calls the Visual Crossing API using the API key.
4. **Caching**: The fetched data is stored in Redis with a time-to-live (TTL) to ensure efficient reuse.
5. **Response**: The data is returned to the client.

---

## Docker and Redis Integration

### Docker Compose File

The `docker-compose.yml` file exists in the project root.

### Key Points:

- The Redis container is exposed on the port specified inside `.env` file.

To stop the Redis container, run:

```bash
docker-compose down
```

---

## Visual Crossing Weather API

- This project uses the **Visual Crossing Weather API** to fetch weather information.
- Visit [Visual Crossing](https://www.visualcrossing.com/) to learn more about their API and obtain an API key.

**API Endpoint** (used internally by the app):

```
https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/{location}?key={WEATHER_API_KEY}
```

---

## Future Scope

While the project is currently designed as a simple learning exercise, it can be extended with additional features:

1. **Support for hourly and weekly forecasts**.
2. **Authentication and rate-limiting** for better security.
3. **Error handling** improvements.
4. **Improved caching strategy** (e.g., different TTLs for different data types).

---

## Summary

This project demonstrates how to:

- Build a Go server that fetches and serves weather data.
- Integrate Redis for caching to improve response times.
- Use Docker Compose to containerize Redis for easy setup and development.
