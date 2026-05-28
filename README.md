# URL Shortener CLI

A lightweight URL shortening service built with a **Node.js + MongoDB backend** and a **Go CLI client**. This project demonstrates API integration, HTTP requests, and command-line interface development.

## 📖 The Story Behind This Project

A few weeks ago, I built the backend for a simple URL Shortener API using Node.js. Later, while learning Go and exploring concepts like HTTP requests, JSON handling, and API communication, I started wondering:

"What if I create a CLI tool in Go to interact with my existing backend?"

That idea quickly turned into this project.

The Go CLI was written from scratch in less than 2 hours and communicates directly with the URL Shortener API. Through this project, I explored:

HTTP GET & POST requests in Go
JSON marshaling and unmarshaling
CLI argument handling
API integration
Backend-client communication
Structuring terminal-based developer tools

What started as a small experiment while learning Go became a really fun way to connect concepts across different technologies.

This project is a reminder that learning a new language becomes much more exciting when you immediately build something practical with it.

## 🏗️ Project Architecture

```
UrlShortenCLI/
├── Backend/                 # Node.js Express Server with MongoDB
│   ├── index.js            # Main server entry point
│   ├── package.json        # Node.js dependencies
│   ├── controllers/        # API request handlers
│   │   └── url.controller.js
│   ├── models/             # MongoDB data schemas
│   │   └── urlShortner.model.js
│   ├── routers/            # API route definitions
│   │   └── url.routes.js
│   └── utils/              # Helper utilities
│       └── hashGenerator.utils.js
└── CLI-client/             # Go Command-Line Interface
    ├── main.go             # CLI application entry point
    └── go.mod              # Go module configuration
```

## 🚀 Getting Started

### Prerequisites

Before running this project, ensure you have installed:

- **Node.js** (v14 or higher)
- **MongoDB** (running locally or remote connection)
- **Go** (v1.26 or higher)
- **npm** (comes with Node.js)

### Backend Setup

1. **Navigate to the Backend folder:**
   ```bash
   cd Backend
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

3. **Configure environment variables:**
   Create a `.env` file in the `Backend/` folder:
   ```
   PORT=8000
   MONGO_URL=mongodb://localhost:27017
   ```
   
   - Replace `mongodb://localhost:27017` with your MongoDB connection string if using a remote database

4. **Start the backend server:**
   ```bash
   npm start
   ```
   
   The server will start on `http://localhost:8000` and will be watching for changes with **nodemon**.

   **Expected output:**
   ```
   Connected to MongoDB successfully.
   Server Started!!
   ```

### CLI Client Setup

1. **Navigate to the CLI-client folder:**
   ```bash
   cd CLI-client
   ```

2. **Build the Go application:**
   ```bash
   go build -o shorten main.go
   ```
   
   This creates an executable named `shorten`

3. **Run the CLI (Development mode):**
   ```bash
   go run main.go <command>
   ```

4. **Or use the compiled binary (Production mode):**
   ```bash
   ./shorten <command>
   ```

## 📝 CLI Commands

### Available Commands

```bash
# Check backend server health
shorten health
shorten --health

# List all shortened URLs created
shorten list
shorten --list

# Create a new shortened URL
shorten create
shorten --create

# View analytics for a specific shortened URL
shorten analyse
shorten --analyse

# Display help documentation
shorten man
shorten --man
```

### Command Examples

**1. Check Server Health:**
```bash
go run main.go health
```
Output:
```
Status:  200
Content length:  24
Res:  Healthy,Up and Running.
```

**2. List All Shortened URLs:**
```bash
go run main.go list
```

**3. Create a Shortened URL:**
```bash
go run main.go create
```
The CLI will prompt you to input a URL:
```
Input a link you want shorten:
https://www.example.com/very/long/url
```

**4. View Analytics:**
```bash
go run main.go analyse
```
You'll be prompted to enter the hash ID of the shortened URL:
```
Input the hash ID assigned:
abc123
```

## 🔌 Backend API Endpoints

### Base URL
```
http://localhost:8000
```

### Health Check
- **Endpoint:** `GET /`
- **Description:** Check if the server is running
- **Response:**
  ```json
  {
    "msg": "Healthy, Up and Running."
  }
  ```

### URL Management

#### Get All URLs
- **Endpoint:** `GET /api/v1/url`
- **Description:** Retrieve all shortened URLs
- **Response:**
  ```json
  {
    "message": "Success",
    "data": [
      {
        "_id": "507f1f77bcf86cd799439011",
        "createdAt": "2026-05-28T10:30:00Z",
        "hash_id": "abc123",
        "original_url": "https://example.com/long/url",
        "short_url": "http://localhost:8000/abc123"
      }
    ]
  }
  ```

#### Create a Short URL
- **Endpoint:** `POST /api/v1/url`
- **Content-Type:** `application/json`
- **Request Body:**
  ```json
  {
    "original_url": "https://example.com/very/long/url"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Success",
    "data": {
      "_id": "507f1f77bcf86cd799439011",
      "hash_id": "abc123",
      "original_url": "https://example.com/very/long/url",
      "short_url": "http://localhost:8000/abc123"
    }
  }
  ```

### Redirects & Analytics

#### Redirect to Original URL
- **Endpoint:** `GET /:code`
- **Description:** Redirects to the original URL and increments click count
- **Example:** `GET /abc123` → Redirects to original URL and tracks the visit

#### Get URL Analytics
- **Endpoint:** `GET /analytics/:code`
- **Description:** Retrieve analytics and metadata for a shortened URL
- **Response:**
  ```json
  {
    "analytics": {
      "_id": "507f1f77bcf86cd799439011",
      "count": 5,
      "ipAddress": ["192.168.1.1", "192.168.1.2"]
    },
    "metadata": {
      "_id": "507f1f77bcf86cd799439011",
      "hash_id": "abc123",
      "original_url": "https://example.com/long/url",
      "short_url": "http://localhost:8000/abc123"
    }
  }
  ```

## 🛠️ Technology Stack

### Backend
- **Framework:** Express.js (v5.2.1)
- **Database:** MongoDB with Mongoose (v9.3.3)
- **Runtime:** Node.js
- **Development Tool:** Nodemon (auto-restart on file changes)
- **Security:** Bcrypt for hashing
- **Config Management:** dotenv

### Frontend/CLI
- **Language:** Go (v1.26.3)
- **Standard Libraries:** net/http, encoding/json, bufio, io, os, strings

## 📦 Project Dependencies

### Backend (package.json)
```json
{
  "bcrypt": "^6.0.0",
  "crypto": "^1.0.1",
  "dotenv": "^17.4.2",
  "express": "^5.2.1",
  "mongoose": "^9.3.3",
  "nodemon": "^3.1.14"
}
```

### CLI Client
Go uses only standard library - no external dependencies needed!

## 📚 Key Features

✅ **Create shortened URLs** - Generate unique hash IDs for long URLs  
✅ **URL Redirection** - Automatically redirect from short to original URL  
✅ **Click Analytics** - Track number of clicks and IP addresses  
✅ **Persistent Storage** - All data stored in MongoDB  
✅ **RESTful API** - Clean, standard REST endpoints  
✅ **CLI Interface** - Easy-to-use command-line tool built in Go  
✅ **Error Handling** - Robust error handling and validation  

## 🔑 Key Components

### Backend Controllers
- **url.controller.js** - Handles URL registration, retrieval, analytics, and redirection logic

### Database Models
- **urlShortner.model.js** - Defines the schema for storing URL data and analytics

### Utilities
- **hashGenerator.utils.js** - Generates unique hash IDs for shortened URLs

### CLI Functions
- `HealthChekupGet()` - Verifies backend availability
- `GetAllShortLinks()` - Fetches all shortened URLs
- `CreateShortLink()` - Prompts user and creates new short URL
- `GetAnalytics()` - Retrieves and displays analytics
- `ManPage()` - Displays help documentation

## 💡 Learning Outcomes

This project demonstrates:
- Building RESTful APIs with Express.js
- Database operations with MongoDB and Mongoose
- HTTP client requests in Go
- Command-line argument parsing
- JSON marshalling/unmarshalling
- Error handling and validation
- Environment configuration management
- Building interactive CLI tools

## 🐛 Troubleshooting

### Backend won't start
- Check if MongoDB is running: `mongosh`
- Verify `.env` file has correct `MONGO_URL`
- Ensure port 8000 is not in use

### CLI not connecting to backend
- Verify backend is running on `http://localhost:8000`
- Check firewall settings
- Ensure backend URL in [CLI-client/main.go](CLI-client/main.go) line 13 matches your setup

### Go build errors
- Update Go: `go version`
- Clear build cache: `go clean`
- Reinstall dependencies: `go mod tidy`

## 📄 License

ISC License

## 👤 Author

**rahulkumarparida**

---

**Built with ❤️ while learning Go and building cool CLI tools!**
