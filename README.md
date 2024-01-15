# OnLend - Financial Management and Payment Backend

## Description

OnLend is a backend application designed for peer-to-peer payment and financial management. It's inspired by the functionality of services like PayPal, with a focus on providing a robust API for managing personal finances, transactions, and more.

- **Motivation**: To create a backend system that simplifies financial transactions and account management for developers building financial applications.
- **Problem Solved**: Offers a comprehensive set of features for handling payments, recurring billing, and splitting expenses among users.
- **Key Learnings**: This project enhanced my skills in RESTful API development, database design, security implementation, and scalable backend architecture.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Credits](#credits)
- [License](#license)
- [How to Contribute](#how-to-contribute)
- [Tests](#tests)

## Installation

1. Clone the repository to your local machine.
2. Ensure Docker and Docker Compose are installed.
3. Run `docker-compose up --build` in the project directory.
4. The API is now accessible at `http://localhost:8081`.

## Usage

Use OnLend's API endpoints for:

- User account creation and management.
- Initiating and receiving peer-to-peer payments.
- Setting up recurring payments and managing subscriptions.
- Splitting bills and tracking shared expenses.


## Features

- **Recurring Payments**: Automate periodic payments and subscriptions.
- **Request Money**: Send payment requests with secure links.
- **Split Bill**: Divide bills among multiple users with ease.
- **Real-time Notifications**: Implement webhook support for transaction notifications.
- **Secure Authentication**: Robust user authentication and authorization.

## Credits

Developed by Marco Braun. Key resources and acknowledgments:

- [Go Programming Language](https://golang.org/)
- [Echo - High performance, extensible, minimalist Go web framework](https://echo.labstack.com/)
- [PostgreSQL](https://www.postgresql.org/)

## License

Distributed under the MIT License. See `LICENSE` for more information.

## How to Contribute

Contributions are welcome! Please follow these steps:

1. Fork the project.
2. Create a feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

## Tests

Run `go test ./...` in the project root to execute the test suite.

## Badges

![Go Report Card](https://goreportcard.com/badge/github.com/user/repo)
![Docker Automated build](https://img.shields.io/docker/automated/user/repo.svg)

---

This README provides a basic overview of the OnLend backend project. Expand and modify as needed to accurately reflect the specifics and progress of your project.

