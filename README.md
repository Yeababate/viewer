# Car Viewer

A Go web application for browsing a car catalog, viewing detailed car pages, and comparing selected vehicles.

## Features

- Browse all cars from a backend API
- Search cars by name or year
- Filter by manufacturer, category, and year
- View paginated results
- Open a dedicated car detail page
- Compare selected cars (exactly 2 at a time)
- Render friendly error pages for missing routes or invalid car IDs

## Tech Stack

- Go
- net/http
- html/template

## Prerequisites

- Go 1.25 or newer
- Node.js and npm (for running the backend API)
- A backend API running on http://localhost:3000

The viewer expects these API endpoints:

- `GET /api/models`
- `GET /api/models/{id}`
- `GET /api/manufacturers/`
- `GET /api/manufacturers/{id}`
- `GET /api/categories`
- `GET /api/categories/{id}`

## Project Layout


```text
project-root/
├── api/
└── viewer/
```

## Running the API

The Cars API is a Node.js/Express app. From the `api` folder:

1. Install required packages:

- Install [NodeJS](http://nodejs.org)
- Install [NPM](https://www.npmjs.com/package/npm) package manager

   ```bash
   make build
   ```

2. Start the server:

   ```bash
   make run
   ```


## Running the Project

1. Make sure the backend API is running (see [Running the API](#running-the-api) above).
2. From the `./viewer` project root, run:

   ```bash
   go run .
   ```

3. Open your browser at: http://localhost:8080

## Project Structure

| File | Purpose |
|---|---|
| `main.go` | Starts the HTTP server and initializes the app |
| `handlers.go` | Route handlers for home, details, and compare pages |
| `getData.go` | Functions for fetching car data from the API |
| `filterAndSearch.go` | Search and filtering logic |
| `pagination.go` | Pagination logic |
| `typesAndVariables.go` | Shared structs and variables |
| `templates/` | HTML templates |
| `static/` | CSS and images |
| `static/img/` | Car images, referenced by templates as `/static/img/{filename}` |



## Notes

- This project is a frontend viewer and depends on an external API for its data. If the API is unavailable, the app will not be able to load the car catalog.

- Car images must be placed in `static/img/` for the viewer to display them.