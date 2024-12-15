# Project Overview

This project consists of two services written in Golang using Domain-Driven Design (DDD). The code structure follows the repository pattern with unit of work and includes an anti-corruption layer for service communication using gRPC.

## Code Architecture

The codebase is structured with a cmd directory containing separate files for each service, allowing the generation of two binary executables. Common components are reused across services, and the repository pattern is implemented to manage data storage.

## Domain and Repositories

The domain layer contains entities and value objects.
There are two repositories: one SQL-based for production use and one in-memory for testing purposes. The in-memory repository speeds up testing by avoiding the need for a database.

## Unit of Work

The Unit of Work pattern is used to abstract over repositories. Each service method runs within a transaction context, ensuring data consistency and integrity.

## Anti-Corruption Layer

An anti-corruption layer is implemented to facilitate communication between services using gRPC. This layer ensures that dependencies are injected into the service, allowing for reliable testing and improved code readability.

## Testing

The testing suite avoids mocking just as how God intended it to be. Each test makes a throwaway container of postgres. This is slow, can made fast using template dbs but I didn't had time.

## Running the Code

There is a docker-compose file, you just need to run:

```
docker-compose up
```

and it will be build everything

there is also a insomnia json exported file that you can import in postman to easily test the service(s)

## What I would have done if I had more time

I was off this week and was home at Saturday, so hard very little time from the start. There are two things I would love to do.

### 1. Event driven Architecture

Go is perfect for this, the idea is to make a event bus which is a channel and that is consumed by different handlers (each of each event). This way everything in the system becomes a event and scales very well.

### 2. Better logging

Currently, there is no logging in the system :(
