# Log ingestor

## Log ingestor

- Accepts logs of format:
  ```
  {
  "level": "error",
  "message": "Failed to connect to DB",
  "resourceId": "server-1234",
  "timestamp": "2023-09-15T08:00:00Z",
  "traceId": "abc-xyz-123",
  "spanId": "span-456",
  "commit": "5e5342f",
  "metadata": {
      "parentResourceId": "server-0987"
  }
  }
  ```
- Logs are ingested via HTTP POST requests to `http://localhost:4000/add` endpoint

### How to setup the dashboard

- Run docker-compose
  ```
  make up
  ```
