# 🛍️ MarketPlace – Scalable Classified Ads Platform in Golang

MarketPlace is a scalable and secure backend system for managing classified ads, built with Go and Gin. It supports OTP-based authentication, ad publishing, fast search, caching, and monitoring. The project also includes a frontend layer, and demonstrates full-stack integration and communication between client and server.

---

## 🚀 Features

- OTP-based authentication with RSA token generation
- Redis-backed OTP storage and token persistence in PostgreSQL
- JWT-based session management
- Rate limiting and blacklist system for abuse prevention
- Ad publishing and public listing
- Full-text search using MongoDB for location-based queries (e.g., "Shiraz")
- Redis caching for fast ad delivery
- Prometheus and Grafana for metrics and query monitoring
- ELK stack (Elasticsearch, Logstash, Kibana) for centralized logging
- Dockerized setup for easy deployment
- Frontend integration with RESTful APIs (React or other frameworks)
- Full understanding of frontend-backend communication and data flow

---

## 📁 Project Structure

MarketPlace/ ├── cmd/ # Application entry point ├── config/ # Configuration files ├── controller/ # HTTP handlers ├── middleware/ # Auth, rate limiting, blacklist ├── model/ # Data models ├── repository/ # PostgreSQL and MongoDB access ├── router/ # Route definitions ├── service/ # Business logic └── utils/ # Utility functions

Code

---

## 🧑‍💻 Getting Started

### Prerequisites:
- Go 1.20+
- PostgreSQL
- MongoDB
- Redis
- Docker

### Run locally:
```bash
git clone https://github.com/sajjadmokhtari/MarketPlace.git
cd MarketPlace
go run cmd/main.go
Run with Docker:
bash
docker-compose up --build
🔐 Authentication Flow
User requests OTP

OTP is stored in Redis

Upon verification, RSA token is generated and stored in PostgreSQL

JWT is issued for session management

Blacklist and rate limiter middleware protect sensitive endpoints

📬 Core API Endpoints
Endpoint	Description
POST /api/otp	Request OTP for login/register
POST /api/verify	Verify OTP and receive token
POST /api/login	Login with token and get JWT
POST /api/ad	Publish a new ad
GET /api/ads	View all published ads
GET /api/search?q=	Search ads by location keyword
🌐 Frontend Integration
The backend is fully integrated with a frontend client (e.g., React). The frontend communicates with the backend via RESTful APIs, handles user sessions with JWT, and displays ads dynamically based on search queries. This project demonstrates full-stack understanding and implementation of client-server architecture.

⚡ Performance & Monitoring
Redis caching for fast ad delivery

MongoDB indexing for efficient search

Prometheus + Grafana dashboards for metrics

ELK stack for centralized log management

🧪 Testing
Unit tests and integration tests are under development using Go's testing and testify.

🤝 Contributing
This project is under active development. Contributions are welcome! Feel free to open an issue or submit a pull request.
