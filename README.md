# Intelligent Document Search for Research Papers (IntelliDoc)

## Project Overview
This project provides an intelligent search interface for research papers using PostgreSQL with Timescale Cloud, leveraging AI-powered embeddings for similarity search.

## Setup Instructions

1. **Environment Variables**:
   - Create a `.env` file with database configuration:
     ```
     DB_HOST=your-timescale-cloud-host
     DB_PORT=your-timescale-cloud-port
     DB_USER=your-username
     DB_PASSWORD=your-password
     DB_NAME=your-database-name
     ```

2. **Running the Project**:
   ```bash
   go run main.go
