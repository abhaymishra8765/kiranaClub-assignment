**Project Documentation for Retail Pulse Backend Service:**

__Description:__

This project is a backend service written in Go, designed to process thousands of images collected from stores. It allows users to submit image processing jobs, where each image's perimeter is calculated. The processing mimics GPU-intensive tasks by adding random delays. Jobs can be submitted via a REST API, and their status can be queried to see if processing is ongoing, completed, or failed.

__Assumptions:__

1.The perimeter is calculated assuming the images have a fixed size (e.g., 100x200) since real image size data is not provided.

2.Random sleep times between 0.1 and 0.4 seconds are used to simulate GPU processing delays.

3.Store Master data is static and loaded from an external file for matching store IDs.



__Installation (Setup) Instructions:__

There are two ways to run the application: via Docker or without Docker.

__With Docker:__

1. Ensure that Docker and Docker Compose are installed.
2. Clone the repository.
3. Build and run the service using Docker Compose:
   
   ![image](https://github.com/user-attachments/assets/48f4c465-a418-43a0-8ac2-7b38d768f941)


   
4.The application will be accessible on port 8080. You can test the API endpoints using tools like Postman or curl.

__Without Docker:__

1.Ensure Go is installed on your system. Install it from here.

2.Clone the repository.

3.Initialize the Go modules and download dependencies


![image](https://github.com/user-attachments/assets/33cfaaa7-d300-453e-b695-d2265687e28b)


4.Run the application

![image](https://github.com/user-attachments/assets/deb51c8a-da19-41b5-b1d8-cb97e3fb36de)

5.The service will run on http://localhost:8080.


__Testing Instructions:__

You can use Postman, curl, or any API testing tool to test the following endpoints.

__Submit a Job:__

Endpoint: POST /api/submit/ 

__Get Job Status:__

Endpoint: GET /api/status?jobid=1


__Work Environment:__

__OS__: Windows 11

__Editor/IDE:__ Visual Studio Code

__Go Version:__ Go 1.23

__Docker Version:__ Docker version 27.3.1

__Libraries/Tools:__

github.com/gorilla/mux for routing.

time and net/http for handling image downloads and processing.

sync and Go’s native concurrency tools for handling multiple jobs.


__Improvements:__

__Scalability:__ Introduce a message queue like RabbitMQ or Kafka to queue and process jobs asynchronously in a more scalable way.

__Error Handling:__ Implement retry logic for image downloads and more comprehensive error handling.

__Job Persistence:__ Use a database like Redis to store job status, results, and logs to ensure persistence across service restarts.

__Image Size Calculation:__ Instead of assuming fixed dimensions, implement logic to determine the actual dimensions of the downloaded images.


