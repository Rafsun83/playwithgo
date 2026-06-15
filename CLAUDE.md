my-backend/
├── cmd/
│ └── server/
│ └── main.go # Application entry point
├── internal/ # Private application code
│ ├── handlers/ # HTTP handlers (controllers)
│ ├── services/ # Business logic
│ ├── repository/ # Database access layer
│ └── models/ # Data structures
├── pkg/ # Public reusable code (optional)
├── config/ # Configuration files
├── migrations/ # Database migrations (if using SQL)
├── .env # Environment variables (don't commit)
├── .gitignore
├── go.mod
├── go.sum
└── README.md
