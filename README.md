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

### Setup Instructions

Follow these steps to set up the development and production environment for HiveBox:

#### 1. **Clone the Repository**

First, clone the repository to your local machine:

`git clone https://github.com/yourusername/hivebox.git`
`cd hivebox`

#### 2. **Install Dependencies**

Make sure you have Go installed on your machine (version 1.16 or higher). Install the required dependencies by running:

`go mod tidy` 

This will install all the Go dependencies needed by the project.

#### 3. **Running the Application Locally**

To run the HiveBox API locally, execute the following command:

`go run main.go` 

This will start the API server locally. You can test the version endpoint by navigating to:


`http://localhost:8080/version` 

For the average temperature endpoint (assuming senseBox data is available):

`http://localhost:8080/temperature`

### Docker Integration

The project includes a `Dockerfile` for containerizing the API. This allows the API to be run in any environment with Docker installed.

#### **Build and Run with Docker**

1.  **Build the Docker Image** From the project root directory, run the following command to build the Docker image:
    
    `docker build -t hivebox:latest .` 
    
2.  **Run the Docker Container** After the image is built, run the container with the following command:
    
    `docker run -d -p 8080:8080 hivebox:latest` 
    
    The API will now be available on port 8080.

### CI/CD Pipeline (GitHub Actions)

This project is configured with a **GitHub Actions** workflow that automates linting, building, and testing.

#### Workflow Overview

The workflow is triggered when there is a **push** or a **pull request** to specific branches (`dev`, `main`, and others). It performs the following jobs:

1.  **Lint**: Runs static analysis using `golangci-lint` to ensure that the code follows consistent coding styles and best practices.
2.  **Build**: Builds the Docker image for the HiveBox API and saves it as an artifact for testing.
3.  **Test**: Runs unit tests using Go’s testing framework and tests the `/version` endpoint by starting the Docker container and verifying its availability.

### Versioning and Tags

This project uses **Semantic Versioning** for release management. Each release is tagged to signify the version, following this structure:

-   **MAJOR**: For breaking changes (incompatible changes to the API).
-   **MINOR**: For new, backward-compatible features.
-   **PATCH**: For small fixes or improvements (non-breaking changes).
