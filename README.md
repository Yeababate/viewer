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
- A backend API running on http://localhost:3000

The viewer expects these API endpoints:

- GET /api/models
- GET /api/models/{id}
- GET /api/manufacturers/
- GET /api/manufacturers/{id}
- GET /api/categories
- GET /api/categories/{id}

## Running the Project

1. Start the backend API on port 3000.
2. From the project root, run:

```bash
go run .
```

3. Open your browser at:

```text
http://localhost:8080
```

## Project Structure

- main.go - starts the HTTP server and initializes the app
- handlers.go - route handlers for home, details, and compare pages
- getData.go - functions for fetching car data from the API
- filterAndSearch.go - search and filtering logic
- pagination.go - pagination logic
- templates/ - HTML templates
- static/ - CSS and images
- typesAndVariables.go - shared structs and variables

## Notes

This project is a frontend viewer and depends on an external API for its data. If the API is unavailable, the app will not be able to load the car catalog.
