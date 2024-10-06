# ChatApp API

Welcome to the **ChatApp API** repository! This is the backend of a chat application built with Go, Gin, GORM, and JWT for authentication. The UI will be implemented in the future. Below you can find details on how to set up and run the project, as well as the available API endpoints.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Configuration](#configuration)
- [Database Migration](#database-migration)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
    - [Register](#register)
    - [Login](#login)
  - [Protected Routes](#protected-routes)
- [Future Plans](#future-plans)
- [Contributing](#contributing)
- [License](#license)

## Features

- User registration
- User login
- JWT-based authentication
- Built using Go, Gin framework, and GORM ORM

## Tech Stack

- **Backend**: Go (Gin Framework)
- **ORM**: GORM
- **Database**: MySQL (you can replace with other databases supported by GORM)
- **Authentication**: JWT (JSON Web Tokens)

## Installation

### Prerequisites

Make sure you have the following installed:

- Go 1.22
- MySQL or other compatible database
- Git

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/chitano/chatapp.git
   cd chatapp
