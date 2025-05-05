# Short URL

A simple URL shortener application built with Go (Gin framework) and SQLite. This application allows users to shorten long URLs and retrieve shortened links. It includes a front-end interface and a back-end API for URL shortening.

## Features

- Shorten long URLs into unique short URLs.
- Redirect users to the original URL using the short URL.
- Front-end interface for submitting URLs and displaying results.
- Copy the shortened URL to the clipboard with a single click.
- Back-end API for URL shortening and redirection.

## Technologies Used

- **Back-end**: Go, Gin framework, SQLite
- **Front-end**: HTML, JavaScript
- **Database**: SQLite for storing URLs and their hashes

## Installation

### Prerequisites

- Go 1.23 or later installed on your machine.
- SQLite installed (optional, as the application uses SQLite by default).

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/short_url.git
   cd short_url
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. Open your browser and navigate to `http://localhost:3010`.

## API Endpoints

### POST `/shorten`

- **Description**: Shortens a given URL.
- **Request Body**:
  ```json
  {
    "url": "https://www.example.com"
  }
  ```
- **Response**:
  ```json
  {
    "hash": "0a137b37"
  }
  ```

### GET `/:hash`

- **Description**: Redirects to the original URL based on the hash.
- **Example**: `http://localhost:3010/0a137b37` redirects to `https://www.example.com`.

## Front-End Usage

1. Open the application in your browser (`http://localhost:3010`).
2. Enter a long URL in the input field and click the "Shorten" button.
3. The shortened URL will be displayed on the screen.
4. Click the "Copy URL" button to copy the shortened URL to your clipboard.

## Project Structure

```
short_url/
├── main.go          # Main application logic
├── templates/
│   └── index.html   # Front-end HTML template
├── test.db          # SQLite database file
├── go.mod           # Go module dependencies
├── LICENSE          # License file
└── README.md        # Project documentation
```

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.