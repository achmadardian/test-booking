# Test Booking to Go: Go + Laravel

This is a full-stack app using:

- **Backend**: Go 1.23
- **Frontend**: Laravel 9 with PHP 8.2
- **Database**: Postgres 17.4
- **Migration**: Go Migrate
- **Deployment**: Docker and Docker Compose

---

## Project Structure

```
test-booking/
├── api/ # Go backend
└── web/ # Laravel frontend
```

---

## Installation

### 1. Clone project

```bash
git clone https://github.com/achmadardian/test-booking.git
```

```bash
cd test-booking
```

### 2. Run Docker Compose

Assume in `root` folder

```bash
docker-compose up --build
```

---

## Access

- Api can be access at: `http://localhost:9090/api`
- Web can be access at: `http://localhost:8080`

---

## API Available


| Endpoint                                | Method | Description                         | Implementation |
| ----------------------------------------| ------ | ------------------------------------|----------------|
| `/api/healthcheck`                      | GET    | Check app health                    |       ✅       |
| `/api/nationalities`                    | GET    | Get all nationalities               |       ✅       |
| `/api/customers`                        | GET    | Get all customers                   |       ✅       |
| `/api/customers`                        | POST   | Create a new customer               |       ✅       |
| `/api/customers/{customer_id}`          | GET    | Get a single customer               |       ✅       |
| `/api/customers/{customer_id}`          | PATCH  | Update a customer                   |       ✅       |
| `/api/customers/{customer_id}`          | DELETE | Delete a customer                   |       ✅       |
| `/api/customers/{customer_id}/families` | GET    | Get a single customer with families |       ✅       |
| `/api/families`                         | GET    | Get all families                    |       ✅       |
| `/api/families`                         | POST   | Create a new familiy                |       ✅       |
| `/api/families/{family_id}`             | GET    | Get a single family                 |       ✅       |
| `/api/families/{family_id}`             | PATCH  | Update a family                     |       ✅       |
| `/api/families/{family_id}`             | DELETE | Delete a family                     |       ✅       |

---

## Web Implementation
| Endpoint                                | Method | Implementation                      |
| ----------------------------------------| ------ | ----------------------------------- |
| `/api/healthcheck`                      | GET    |                  ❌                 |
| `/api/nationalities`                    | GET    |                  ❌                 |
| `/api/customers`                        | GET    |                  ✅                 |
| `/api/customers`                        | POST   |                  ❌                 |
| `/api/customers/{customer_id}`          | GET    |                  ✅                 |
| `/api/customers/{customer_id}`          | PATCH  |                  ❌                 |
| `/api/customers/{customer_id}`          | DELETE |                  ✅                 |
| `/api/customers/{customer_id}/families` | GET    |                  ✅                 |
| `/api/families`                         | GET    |                  ❌                 |
| `/api/families`                         | POST   |                  ❌                 |
| `/api/families/{family_id}`             | GET    |                  ❌                 | 
| `/api/families/{family_id}`             | PATCH  |                  ✅                 |
| `/api/families/{family_id}`             | DELETE |                  ❌                 |