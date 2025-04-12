# My Go Project

This is a simple Go project that demonstrates the structure of a Go application with a command-line interface and utility functions.

## Project Structure

```sh
my-go-project
├── cmd
│   └── main.go        # Entry point of the application
├── pkg
│   └── utils
│       └── helper.go  # Utility functions
├── go.mod             # Module definition
└── README.md          # Project documentation
```

## Getting Started

To get started with this project, follow the instructions below.

### Prerequisites

Make sure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

### Building the Project

To build the project, navigate to the project directory and run:

```sh
go build -o my-go-app ./cmd
```

### Running the Application

After building, you can run the application with:

```sh
./my-go-app 1 1
```

### Usage

This project currently includes utility functions for basic arithmetic operations. You can extend the functionality by adding more utilities in the `pkg/utils` directory.

#### Locally

```sh
sh run-local.sh
```

