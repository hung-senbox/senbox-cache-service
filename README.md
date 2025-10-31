# Senbox Cache Service

A Go-based caching service built on Redis for managing profile code caching operations. This service provides efficient caching and retrieval of codes for various user types (users, students, teachers, staff, parents, and children).

## Overview

This cache service provides a clean interface for caching and retrieving profile codes with automatic expiration (TTL). It implements a two-layer approach:
- **Caching Layer**: For setting and invalidating cached values
- **Cached Layer**: For retrieving cached values

## Features

- ✅ Redis-based caching with configurable TTL
- ✅ Support for multiple profile types (User, Student, Teacher, Staff, Parent, Child)
- ✅ Automatic cache expiration
- ✅ Type-safe cache key generation
- ✅ Configurable Redis connection via YAML
- ✅ Clean interface-based design

## Requirements

- Go 1.24.4 or higher
- Redis server

## Installation

```bash
go mod download
```

## Configuration

Create a YAML configuration file (e.g., `config.yaml`) with the following structure:

```yaml
redis_cache:
  host: localhost
  port: 6379
  password: ""
  db: 0
  ttlSeconds: 3600
```

### Configuration Parameters

- `host`: Redis server hostname or IP address
- `port`: Redis server port
- `password`: Redis authentication password (leave empty if not required)
- `db`: Redis database number
- `ttlSeconds`: Default time-to-live in seconds for cached values

## Usage

### Initializing Redis Cache

```go
import (
    "github.com/hung-senbox/cache-service/pkg/redis"
)

// Initialize Redis cache from config file
cache, err := redis.InitRedisCacheFromFile("config.yaml")
if err != nil {
    log.Fatal(err)
}
```

### Setting Cache Values (Caching Service)

```go
import (
    "github.com/hung-senbox/cache-service/pkg/cache"
    "github.com/hung-senbox/cache-service/pkg/cache/caching"
)

// Create caching service with default TTL
cachingService := caching.NewCachingService(cache, 3600) // TTL: 1 hour

// Set codes for different profile types
ctx := context.Background()

// Set user code
err := cachingService.SetUserCode(ctx, "user123", "ABC123")

// Set student code
err = cachingService.SetStudentCode(ctx, "student456", "DEF456")

// Set teacher code
err = cachingService.SetTeacherCode(ctx, "teacher789", "GHI789")

// Invalidate cache
err = cachingService.InvalidateUserCode(ctx, "user123")
```

### Retrieving Cache Values (Cached Gateway)

```go
import (
    "github.com/hung-senbox/cache-service/pkg/cache/cached"
)

// Create cached gateway
cachedGateway := cached.NewCachedProfileGateway(cache, 3600)

// Retrieve codes
ctx := context.Background()

userCode, err := cachedGateway.GetUserCode(ctx, "user123")
studentCode, err := cachedGateway.GetStudentCode(ctx, "student456")
teacherCode, err := cachedGateway.GetTeacherCode(ctx, "teacher789")
```

### Available Methods

#### Caching Service (Set/Invalidate)
- `SetUserCode(ctx, userID, code)` - Cache user code
- `SetStudentCode(ctx, studentID, code)` - Cache student code
- `SetTeacherCode(ctx, teacherID, code)` - Cache teacher code
- `SetStaffCode(ctx, staffID, code)` - Cache staff code
- `SetParentCode(ctx, parentID, code)` - Cache parent code
- `SetChildCode(ctx, childID, code)` - Cache child code
- `InvalidateUserCode(ctx, userID)` - Remove cached user code
- `InvalidateStudentCode(ctx, studentID)` - Remove cached student code
- `InvalidateTeacherCode(ctx, teacherID)` - Remove cached teacher code
- `InvalidateStaffCode(ctx, staffID)` - Remove cached staff code
- `InvalidateParentCode(ctx, parentID)` - Remove cached parent code
- `InvalidateChildCode(ctx, childID)` - Remove cached child code

#### Cached Gateway (Get)
- `GetUserCode(ctx, userID)` - Retrieve cached user code
- `GetStudentCode(ctx, studentID)` - Retrieve cached student code
- `GetTeacherCode(ctx, teacherID)` - Retrieve cached teacher code
- `GetStaffCode(ctx, staffID)` - Retrieve cached staff code
- `GetParentCode(ctx, parentID)` - Retrieve cached parent code
- `GetChildCode(ctx, childID)` - Retrieve cached child code

## Cache Key Format

Cache keys follow a consistent naming pattern:
- User: `profile-service:user_code:{userID}`
- Student: `profile-service:student_code:{studentID}`
- Teacher: `profile-service:teacher_code:{teacherID}`
- Staff: `profile-service:staff_code:{staffID}`
- Parent: `profile-service:parent_code:{parentID}`
- Child: `profile-service:child_code:{childID}`

## Project Structure

```
senbox-cache-service/
├── pkg/
│   ├── cache/
│   │   ├── cache.go                 # Cache interface definition
│   │   ├── redis_cache.go           # Redis implementation
│   │   ├── caching/
│   │   │   └── caching_profile_service.go  # Service for setting cache
│   │   ├── cached/
│   │   │   └── cached_profile_service.go   # Gateway for getting cache
│   │   └── keys_cache/
│   │       └── profile_keys_cache.go      # Cache key generators
│   ├── config/
│   │   └── config.go                # Configuration loader
│   └── redis/
│       └── redis.go                 # Redis initialization
└── helper/
    └── constants.go                 # Cache prefix constants
```

## Dependencies

- `github.com/redis/go-redis/v9` - Redis client library
- `github.com/gin-gonic/gin` - HTTP web framework
- `gopkg.in/yaml.v3` - YAML configuration parsing

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build
```

## License

This project is part of the Senbox ecosystem.
