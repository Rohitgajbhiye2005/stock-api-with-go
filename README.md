# 📈 Stock API with Go (Golang + PostgreSQL)

A simple RESTful API built using **Go (Golang)** and **PostgreSQL** for managing stocks. This project demonstrates basic CRUD operations using `net/http`, `gorilla/mux`, and `database/sql`.

---

## 🚀 Features

- 📥 Create a new stock
- 📄 Read all stocks or a specific stock by ID
- ✏️ Update an existing stock
- ❌ Delete a stock
- 🔐 Environment-based DB config using `.env`

---

## 🛠 Tech Stack

- **Backend:** Go (Golang)
- **Database:** PostgreSQL
- **Router:** Gorilla Mux
- **Env Loader:** godotenv

---

## 📁 Folder Structure

stock-api/ │ ├── middleware/ # Database logic (CRUD functions) ├── main.go # Main application entry ├── go.mod # Module file ├── .env # Environment config (not pushed) └── README.md

yaml
Copy
Edit

---

## ⚙️ Setup Instructions

1. **Clone the repo:**
   ```bash
   git clone https://github.com/Rohitgajbhiye2005/stock-api-with-go.git
   cd stock-api-with-go
Create .env file:

env
Copy
Edit
POSTGRES_URL=postgres://<username>:<password>@localhost:5432/stocksdb?sslmode=disable
Install dependencies:

bash
Copy
Edit
go mod tidy
Run the app:

bash
Copy
Edit
go run main.go
📬 API Endpoints
Method	Endpoint	Description
POST	/api/newstock	Create a new stock
GET	/api/stocks	Get all stocks
GET	/api/stock/{id}	Get stock by ID
PUT	/api/stock/{id}	Update stock by ID
DELETE	/api/stock/{id}	Delete stock by ID
📦 Sample Request (POST)
URL: http://localhost:8080/api/newstock
Headers:

http
Copy
Edit
Content-Type: application/json
Body:

json
Copy
Edit
{
  "name": "Apple",
  "price": 180
}
✅ TODO (Optional Improvements)
Add Swagger/OpenAPI docs

Add JWT-based authentication

Add unit tests

Dockerize the app

🙌 Author
Rohit Gajbhiye
GitHub: @Rohitgajbhiye2005

📄 License
This project is open source and available under the MIT License.
