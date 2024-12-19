# Stress Test CLI

A command-line tool written in Go for performing load tests on web services. It helps developers and QA teams evaluate the performance of APIs and HTTP services by simulating multiple concurrent requests and generating detailed reports.

## Features

- Simulates load with adjustable concurrency levels.
- Generates comprehensive reports, including:
  - Total execution time.
  - Total number of requests performed.
  - Count of requests with HTTP 200 status.
  - Distribution of other HTTP status codes.
- Easy execution via CLI or Docker.

## Installation

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/stress-test-cli.git
   cd stress-test-cli
   ```

2. Build the application:
   ```bash
   go build -o stress-test-cli .
   ```

3. Run the application:
   ```bash
   ./stress-test-cli --url=http://example.com --requests=100 --concurrency=10
   ```

### Using Docker

1. Build the Docker image:
   ```bash
   docker build -t stress-test-cli-app .
   ```

2. Run the application:
   ```bash
   docker run stress-test-cli-app --url=http://example.com --requests=100 --concurrency=10
   ```

## Usage

### CLI Arguments

| Argument      | Description                          | Example                          |
|---------------|--------------------------------------|----------------------------------|
| `--url`       | The URL of the service to test       | `--url=http://example.com`      |
| `--requests`  | Total number of requests to perform | `--requests=100`                |
| `--concurrency` | Number of simultaneous requests     | `--concurrency=10`              |

### Example

```bash
./stress-test-cli --url=http://google.com --requests=10 --concurrency=2
```

### Sample Output

```
================= Report =================
Total execution time: 5s
Total requests performed: 10
Total requests with status 200: 8
Status code distribution:
  200: 8
  429: 2
=========================================
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for improvements or bug fixes.