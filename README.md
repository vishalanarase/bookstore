# Bookstore API

> RESTful API for managing bookstore

## Structure

Overall directory structure:

* `api`: This directory contains files related to the API functionalities, including endpoints, controllers, and API specific logic
* `bin`: This directory stores executable files after compilation
* `build`: This directory houses build-related files e.g. Dockerfile
* `cmd`: This directory contains command-line interface (CLI) and application program interface (API) or executable application entry points
* `docs`: This directory is dedicated to project documentation, including swagger, http request
* `internal`: This directory contains modules meant to be used within the project
* `migrations`: This directory contains database migration scripts or files used for managing changes to the project's database schema
* `tests`: This directory encompasses files related to testing the project, including integration tests

This organized structure aims to provide clarity and maintainability, making it easier to locate, modify, and expand upon different aspects of the project. Please refer to individual directories for more specific details on their contents and functionalities.

### Libraries/Tools used for this project

* `gin`: Gin is a high-performance HTTP web framework for building APIs in Go. It provides routing, middleware support, and a robust set of features for creating web applications efficiently.
* `mysql`: MySQL is a widely used relational database management system. In this project, the MySQL database is utilized, likely via the gorm library for database interactions.
* `gorm`: GORM is an Object-Relational Mapping (ORM) library for Go, designed to simplify database interactions. It provides a convenient way to work with databases by abstracting the underlying SQL queries, enabling developers to work with Go structs directly.
* `soda`: Soda is a command-line tool for handling database migrations in Go projects. It offers functionalities to manage database schemas, versioning, and applying migrations easily.
* `logrus`: Logrus is a structured logging library for Go. It allows flexible logging with various log levels, formatting options, and hooks, making it easier to handle and manage logs in the application.
* `mock`: The mock library is commonly used for creating mock implementations in Go tests. It helps in generating mock objects or interfaces to simulate behavior for testing purposes, especially useful when working with interfaces or external dependencies.

These libraries and tools are integral to the project, providing essential functionalities for web development, database management, logging, testing, and more. They enhance productivity, maintainability, and the overall quality of the GoLang project.
