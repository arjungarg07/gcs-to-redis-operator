# GCS-Redis Operator

## Overview
The GCS-Redis Operator is a Go-based application designed to perform operations on a Redis instance using data from Google Cloud Storage (GCS). It supports batch operations, such as setting multiple keys in Redis (`mset`), and allows configuration of batch sizes and key expirations. The operator efficiently handles data transfer and processing by utilizing channels and asynchronous operations.

## Features
- **GCS to Redis Data Transfer**: Reads data from a specified GCS path and processes it into a Redis instance.
- **Configurable Operations**: Supports different Redis operations like `mset`.
- **Batch Processing**: Allows configuration of batch sizes for efficient data processing.
- **Key Expiration**: Configurable expiration time for Redis keys.

## Getting Started

### Prerequisites
- Go (Golang) installed on your machine.
- Access to a GCS bucket.
- A Redis instance accessible via network.

### Installation
1. **Clone the Repository:**
   ```sh
   git clone <repository-url>
   cd gcs-redis-operator
   ```

2. **Install Dependencies:**
   Make sure all the necessary dependencies are installed:
   ```sh
   go mod tidy
   ```

### Configuration
Before running the application, you need to set up the configuration:

1. **Config Initialization:**
   The configuration is initialized using the `config.InitConfig()` method. Ensure that your configuration files or environment variables are properly set.

### Running the Application
Run the application using the following command with appropriate flags:

```sh
go run main.go --gcsPath="gs://your-bucket/path/to/object" --redisInstance="your-redis-instance-id" --operation="mset" --expiry="3600" --batchSize="100"
```

#### Command-line Flags:
- `--gcsPath`: The GCS path to the data (required).
- `--redisInstance`: The Redis instance ID (required).
- `--operation`: The Redis operation to perform (e.g., `mset`) (required).
- `--expiry`: The expiration time for Redis keys in seconds (required).
- `--batchSize`: The size of the batch for processing data (default: as per constants).

### Example
```sh
go run main.go --gcsPath="gs://example-bucket/data.json" --redisInstance="redis-instance-1" --operation="mset" --expiry="600" --batchSize="50"
```

## Development

### Project Structure
- `config`: Contains configuration setup.
- `constants`: Holds constant values used across the application.
- `dto`: Data Transfer Objects definitions.
- `redis`: Contains Redis client and operations logic.
- `utils`: Utility functions, including GCS file handling.

### Dependencies
- [sonic](https://github.com/bytedance/sonic): High-performance JSON library for Go.
- [flag](https://pkg.go.dev/flag): Package for parsing command-line flags.

## Contribution
Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Acknowledgments
- The developers and maintainers of the libraries and tools used in this project.
- The Go community for providing excellent resources and support.
