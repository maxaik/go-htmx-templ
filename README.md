# Go HTMX Echo Templ Project

This project is a web application built with [Go](https://golang.org/), [HTMX](https://htmx.org/), [Echo](https://echo.labstack.com/), and templated HTML. It uses [Tailwind CSS](https://tailwindcss.com/) for styling and [Air](https://github.com/cosmtrek/air) for live reloading during development.

## Project Structure

- **Go**: Handles server-side logic with Echo as the web framework.
- **HTMX**: Enables dynamic HTML updates without JavaScript.
- **Echo**: Fast and minimalist web framework for Go.
- **Templ**: A powerful and flexible templating engine for Go, offering fast performance and intuitive syntax.
- **Tailwind CSS**: Utility-first CSS framework for styling.
- **Air**: Provides live reloading for Go development, automatically rebuilding and restarting the server when changes are detected.

## Prerequisites

Before running the project, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.16+ recommended)
- [Air](https://github.com/cosmtrek/air) (for live reloading)
- [Tailwind CSS CLI](https://tailwindcss.com/docs/installation)

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-htmx-echo-templ.git
cd go-htmx-echo-templ
```

### 2. Start tailwind watcher
```bash
tailwindcss -i css/input.css -o css/output.css --watch
```

### 3. Start live-reload with air
```bash
air
```
