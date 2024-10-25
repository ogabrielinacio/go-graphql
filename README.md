
# Go GraphQL API

A Golang API that utilizes GraphQL for managing courses and categories. This project provides a simple interface for querying and mutating data related to courses and categories using GraphQL.

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go

### Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/ogabrielinacio/go-graphql.git
   cd go-graphql
   ```

2. **Start the database:**

   You can start the PostgreSQL database using Docker Compose. Make sure you have the required environment variables set in the `envs` file.

   ```bash
   docker-compose --env-file envs up
   ```

3. **Create the necessary tables:**

   Once the PostgreSQL container is up and running, connect to the database and create the following tables:

   ```sql
   CREATE TABLE Categories (
     Id VARCHAR(255),
     Name VARCHAR(255),
     Description TEXT
   );

   CREATE TABLE Course (
     Id VARCHAR(255),
     Name VARCHAR(255),
     Description TEXT,
     CategoryId VARCHAR(255)
   );
   ```

4. **Set environment variables:**

   Ensure you have the correct environment variables set in the `envs` by running the following command:

   ```bash
   export $(grep -v '^#' envs | xargs)
   ```

5. **Run the application:**

   Finally, start the Golang application by running the following command:

   ```bash
   go run cmd/server/server.go
   ```

### Usage

Once the application is running, you can interact with the GraphQL API by sending queries and mutations to manage `courses` and `categories`.

Now you can execute the queries and mutations in the `queries.graphql` file.