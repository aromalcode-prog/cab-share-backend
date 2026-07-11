# Cab Share Backend

## 1. Project Vision

Build a backend service that enables people traveling in similar directions to share rides with each other.

The primary objective of this project is to learn and demonstrate production-level backend development concepts while building a product that could realistically evolve into a real-world application.

The backend should be designed with scalability, maintainability and clean architecture in mind rather than simply implementing CRUD APIs.

---

# 2. Problem Statement

Daily commuting is expensive and inefficient.

Many people travel between similar locations every day with empty seats in their vehicles while others book expensive cabs.

This application connects drivers and passengers travelling in the same direction, allowing them to share rides and split travel costs.

---

# 3. Goals

## Functional Goals

- User Registration
- User Login
- Offer a Ride
- Search Available Rides
- Book a Ride
- Cancel Booking

## Technical Goals

- Learn Clean Architecture
- Learn REST API Design
- Learn Authentication & Authorization
- Learn Database Design
- Learn PostgreSQL
- Learn GORM
- Learn Docker
- Learn Repository Pattern
- Learn Service Layer
- Learn Production Project Structure

---

# 4. MVP Scope

Version 1 focuses only on the essential ride-sharing workflow.

A user should be able to

1. Register
2. Login
3. Offer a ride
4. Search rides
5. View ride details
6. Book a ride
7. Cancel a booking

Features intentionally excluded from MVP

- Maps
- Live location
- Payments
- Notifications
- Ratings
- Chat
- OTP Login
- Google Login

These will be added in future iterations.

---

# 5. User Journey

### Driver

Register
↓

Login
↓

Offer Ride
↓

Passengers discover ride
↓

Passengers book seats
↓

Ride completed

---

### Passenger

Register
↓

Login
↓

Open Home Page
↓

Current location detected automatically

↓

Enter destination

↓

Search rides

↓

Select ride

↓

Book seat

---

# 6. Functional Requirements

## User

- Register account
- Login
- View profile
- Update profile

---

## Ride

- Create ride
- Search rides
- View ride details
- Cancel ride
- Update ride

---

## Booking

- Book available ride
- Cancel booking
- View booking history

---

# 7. Non Functional Requirements

- Clean Architecture
- Modular codebase
- Scalable project structure
- Environment-based configuration
- Dockerized database
- Secure password storage (bcrypt)
- JWT Authentication
- Proper error handling
- Logging
- Easy deployment

---

# 8. Database Design

## User

| Field | Type |
|-------|------|
| ID | uint |
| Name | string |
| Email | string |
| Phone | string |
| PasswordHash | string |
| CreatedAt | time.Time |
| UpdatedAt | time.Time |

---

## Ride

| Field | Type |
|-------|------|
| ID | uint |
| DriverID | uint |
| Source | string |
| Destination | string |
| DepartureTime | time.Time |
| AvailableSeats | int |
| PricePerSeat | float |
| Status | string |
| CreatedAt | time.Time |
| UpdatedAt | time.Time |

---

## Booking

| Field | Type |
|-------|------|
| ID | uint |
| RideID | uint |
| PassengerID | uint |
| SeatsBooked | int |
| Status | string |
| CreatedAt | time.Time |
| UpdatedAt | time.Time |

---

# 9. Entity Relationships

User

1 ---- N Ride

One driver can create multiple rides.

Ride

1 ---- N Booking

One ride can have multiple bookings.

User

1 ---- N Booking

One passenger can make multiple bookings.

---

# 10. Authentication

MVP

Email + Password

Passwords will never be stored directly.

Passwords will be hashed using bcrypt before storing in PostgreSQL.

Future

- Phone OTP
- Google Login
- Apple Login

---

# 11. API Design

## Authentication

POST /register

POST /login

---

## User

GET /profile

PUT /profile

---

## Ride

POST /rides

GET /rides

GET /rides/:id

PUT /rides/:id

DELETE /rides/:id

---

## Booking

POST /rides/:id/book

DELETE /bookings/:id

GET /bookings

---

# 12. Project Structure

cmd/
    api/

config/

internal/
    database/
    models/
    handlers/
    services/
    repositories/
    middleware/

docs/

---

# 13. Development Roadmap

## Phase 1

- Project Bootstrap
- Docker
- PostgreSQL
- GORM
- User Model

---

## Phase 2

Authentication

- Register
- Login
- JWT
- Password Hashing

---

## Phase 3

Ride Management

- Create Ride
- Search Ride
- Ride Details
- Update Ride

---

## Phase 4

Booking

- Book Ride
- Cancel Booking
- Booking History

---

## Phase 5

Production Improvements

- Pagination
- Validation
- Logging
- Middleware
- Unit Testing

---

## Phase 6

Advanced Features

- Maps Integration
- Nearby Ride Search
- Notifications
- Ratings
- Reviews
- Payments
- Admin Dashboard

# Engineering Principles

- Follow Clean Architecture.
- Every package should have a single responsibility.
- Business logic should never live inside handlers.
- Database queries should only exist in repositories.
- Services should contain business logic.
- Prefer composition over global state.
- Never hardcode secrets.
- Use environment variables for configuration.
- Always return meaningful errors.
- Follow Go conventions and idiomatic code.
- Keep functions small and focused.
- Write code for readability first, cleverness second.