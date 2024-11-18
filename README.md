# HiveBox API

This project is an API that provides application version information and average temperature based on senseBox data. It is implemented in Go using the Gin framework to create RESTful endpoints.


# Description
HiveBox is an API designed to obtain two types of information:

1. Application version.
2. Average temperature based on senseBox data.

This API follows Semantic Versioning (v0.0.1) and ensures that the temperature is calculated based on the most recent data, not older than 1 hour.

# Requirements
- Go (version 1.16 or higher).
- Gin (Go Framework to build the API).
- Test suite: Testify (to perform unit tests).
