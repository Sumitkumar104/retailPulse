# Retail Pulse Job Processing Service

## Description

This project provides a backend service that processes image data from various retail stores. It simulates image perimeter calculations (2 * [Height + Width]) and imitates GPU processing by adding random sleep times. The service allows users to submit jobs with store information and image URLs, processes the images, and stores the results.

### Key Features:
- Submit jobs containing store information and image URLs.
- Process images and calculate their perimeter.
- Simulate GPU processing with random delays.
- Handle errors related to missing stores or image download failures.
- Track the status of submitted jobs (ongoing, completed, or failed).
- Store processed results and errors in a structured format.

## Assumptions

- The store data is assumed to be available in a predefined "Store Master" dataset.
- Each image URL provided in the request should point to an accessible image.
- The store ID must exist in the "Store Master" dataset. If not, the job will fail for that store.
- The job processing may take some time depending on the number of images and the complexity of the data.
- A valid job ID must be provided to track the status of a job.

## Work Environment

The following environment was used to develop and run this project:

### 1. Operating System
- **macOS** (Version 12.1 Monterey)  
  Alternatively, this project should run smoothly on **Linux** or **Windows** with appropriate setup.

### 2. IDE/Text Editor
- **Visual Studio Code (VS Code)** (Version 1.60.0)  
  This IDE was used for writing and debugging Go code. It's lightweight, feature-rich, and has excellent support for Go development with extensions like "Go" by Microsoft.

- Alternatively, **GoLand** or **IntelliJ IDEA** with Go plugin can also be used for a more integrated Go development experience.

### 3. Programming Language
- **Go (Golang)** (Version 1.18 or later)  
  The project was written in Go, a statically typed, compiled language known for its simplicity, performance, and ease of concurrency. Go is well-suited for building backend services and APIs.

### 4. Libraries/Dependencies
- **Standard Go Libraries**:  
  These include packages for handling HTTP requests (`net/http`), encoding/decoding JSON (`encoding/json`), managing concurrency (`sync`), working with dates and times (`time`), and more.
  
- **Image Processing**:  
  The project uses the `image` package to decode image files and perform calculations. The packages `image/jpeg` and `image/png` are used to handle specific image formats.

- **Third-Party Libraries**:  
  None for this project specifically, as it was built using Go's standard libraries, but additional libraries like `gorilla/mux` for routing could be integrated if the project scales further.

This project should be compatible across multiple platforms as long as Go and the necessary dependencies are set up correctly.


## Improvements if Given More Time

- Database Integration: Implement persistent storage for jobs, store data, and image processing results in a relational or NoSQL database (e.g., PostgreSQL or MongoDB) for better scalability and long-term data storage.

- Error Handling: Improve error handling with more detailed and specific error responses (e.g., different error codes for different types of failures such as store not found, image processing error, etc.).

- Image Storage: Integrate cloud storage (e.g., Amazon S3, Google Cloud Storage) for storing images, which will allow handling large volumes of images more efficiently and facilitate easier image retrieval.

- API Documentation: Use API documentation tools like Swagger to automatically generate and maintain up-to-date API documentation, making it easier for users to interact with the API.

- Unit and Integration Testing: Write unit and integration tests using testing libraries like testing in Go to ensure the stability and correctness of the application. This will help in reducing bugs and improving the maintainability of the codebase.

- Rate Limiting: Implement rate limiting on the API endpoints to avoid overloading the system with too many requests in a short time.

- Job Queue: Introduce a job queue mechanism (e.g., using Redis or RabbitMQ) to queue jobs for processing and manage the job lifecycle more efficiently, allowing for better load distribution and control.

- User Authentication and Authorization: Add user authentication (e.g., using OAuth or JWT) and role-based access control (RBAC) to secure the API and allow for user-specific job tracking and management.
