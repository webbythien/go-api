# How to Use Go API template


## Requirements

To run Go API, you need the following:

- Golang 1.16 or higher
- Running RabbitMQ environment
- Running Redis environment
- Running mongodb environment

## Installation

1. Clone the source code from the repository:


2. Install the required dependencies
```bash
go mod download
```

3. Run the following command
- run api server
```bash
go run main.go server
```
- run worker server
```bash
go run main.go worker default_queue
```
- options:

```bash
go run main.go worker [queue-name]
```
