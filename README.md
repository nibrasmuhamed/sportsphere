# Sport Sphere Project

**Sport Sphere** is a **multi-tenant turf booking system** built with **Golang** and **MongoDB**. This project is designed to provide efficient and scalable management of turf bookings.

---

## Features

-   **Multi-Tenant Support**: Manage multiple organizations seamlessly within a single system.
-   **Fast Backend**: Built using Golang for high performance and scalability.
-   **MongoDB Replica Set**: Ensures high availability and data reliability.

---

## Prerequisites

1. **Go** (Version 1.18 or higher is recommended).
2. **MongoDB** with Replica Set enabled.

---

## How to Run

## Using Docker

```bash
docker-compose up
```

## Using Local Machine

### 1. Start MongoDB with Replica Set

Run MongoDB in **replica set mode**:

```bash
sudo mongod --replSet rs0 --bind_ip_all --port 27017
```

### 2. Environment Variables

```bash
export CONFIG_PATH=config/config.json
export MONGO_URI="mongodb://localhost:27017"
```

### 3. Update config.json (Optional)

configs can be updated from config/config.json

### 4. Run the Application

```bash
go run main.go
```
