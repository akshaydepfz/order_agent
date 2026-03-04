# Order Agent

A Go-based order management system with WhatsApp integration and AI-powered order processing.

## Structure

```
├── internal/
│   ├── config/       # Configuration loading
│   ├── database/     # PostgreSQL connection
│   ├── middleware/   # Auth middleware
│   ├── models/       # Domain models
│   ├── repository/   # Data access layer
│   ├── services/     # Business logic
│   ├── handlers/     # HTTP handlers
│   └── routes/       # Route setup
├── pkg/utils/        # Shared utilities (JWT, response helpers)
├── migrations/       # SQL migrations
└── main.go
```

## Setup

1. Copy `.env.example` to `.env` and fill in your values.
2. Create a PostgreSQL database and run migrations:
   ```bash
   psql -d your_db -f migrations/001_init.sql
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

## API Endpoints

- `GET /` - Health check
- `POST /auth/login` - Login
- `GET /webhook/whatsapp` - WhatsApp webhook verification
- `POST /webhook/whatsapp` - WhatsApp webhook (incoming messages)
- `GET /api/admin/dashboard` - Admin dashboard (auth required)
- `POST /api/shops` - Create shop (auth required)
- `GET /api/shops/:id` - Get shop by ID (auth required)

## Environment Variables

| Variable | Description |
|----------|-------------|
| PORT | Server port (default: 8080) |
| DATABASE_URL | PostgreSQL connection string |
| JWT_SECRET | Secret for JWT signing |
| WHATSAPP_TOKEN | WhatsApp Business API token |
# order_agent
# order_agent
