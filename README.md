# URL Shortener

A simple URL shortener app built with **Go** (Gin) for the backend and **HTML, CSS (Bulma), JavaScript** for the frontend. It allows users to shorten URLs, view, and delete them.

## Features
- Shorten URLs
- View and delete user-specific URLs
- Responsive frontend (Bulma)

## Setup

### Backend (Go)
1. Install Go: https://golang.org/dl/
2. Clone the repository:
   ```bash
   git clone https://github.com/AndreaPallotta/url_shortener
   cd url_shortener
   ```
3. Install Go modules:
   ```bash
   go mod tidy
   ```
4. Run the server:
   ```bash
   go run main.go
   ```
   Backend runs on port `8080`.

### Frontend (HTML, CSS, JS)
1. Install Node.js: https://nodejs.org/
2. Navigate to the frontend directory:
   ```bash
   cd <frontend-directory>
   ```
3. Install dependencies:
   ```bash
   npm install
   ```
4. Start the frontend:
   ```bash
   npm start
   ```
   Frontend runs on port `8081`.

## Usage
- **Backend**: [http://localhost:8080](http://localhost:8080)
- **Frontend**: [http://localhost:8081](http://localhost:8081)

## License
MIT License
