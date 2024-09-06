# Portgen

Portgen is a CLI tool designed to generate random unused ports within a specified range. It was born out of the frustration of having to think of unused ports for new projects, especially when dealing with multiple running containers.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install Portgen, you need to have Go installed on your system. Then, you can use the following command:

```bash
go install github.com/Lutefd/portgen@latest
```

Alternatively, you can clone the repository and build it manually:

```bash
git clone https://github.com/Lutefd/portgen.git
cd portgen
make build
```

## Usage

Portgen can be used in both interactive and non-interactive modes.

### Non-interactive Mode

```bash
portgen [flags]
```

Flags:
- `-m, --min int`: Minimum port number (inclusive) (default 10000)
- `-M, --max int`: Maximum port number (inclusive) (default 65535)
- `-c, --copy`: Copy the generated port to clipboard
- `-s, --short`: Print only the generated port number
- `-h, --help`: Display help for portgen

### Interactive Mode

To enter interactive mode, simply run `portgen` without the `-s` flag. In this mode, you can use the following commands:

- `generate`: Generate a new port
- `copy`: Copy the current port to clipboard
- `help`: Show the help message

## Project Structure

The project is organized as follows:

```
portgen/
├── cmd/
│   └── portgen/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── app.go
│   │   └── app_test.go
│   ├── cli/
│   │   ├── cli.go
│   │   └── cli_test.go
│   ├── port/
│   │   ├── port.go
│   │   └── port_test.go
│   └── ui/
│       ├── model.go
│       ├── model_test.go
│       ├── styles.go
│       └── styles_test.go
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```

- `cmd/portgen`: Contains the main entry point of the application.
- `internal/app`: Implements the core functionality of generating ports and clipboard operations.
- `internal/cli`: Handles the command-line interface using Cobra.
- `internal/port`: Manages port generation and checking if a port is in use.
- `internal/ui`: Implements the interactive UI using Bubble Tea.

## Development

To set up the development environment, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/Lutefd/portgen.git
   cd portgen
   ```

2. Install dependencies:
   ```bash
   make deps
   ```

3. Build the project:
   ```bash
   make build
   ```

4. Run tests:
   ```bash
   make test
   ```

### Makefile Commands

The project includes a Makefile with the following commands:

- `make all`: Run tests and build the project
- `make build`: Build the project
- `make test`: Run all tests
- `make clean`: Clean up build artifacts
- `make run`: Build and run the project
- `make deps`: Install dependencies
- `make build-linux`: Build for Linux (cross-compilation)

## Contributing

Contributions to Portgen are welcome! Here are some ways you can contribute:

1. Report bugs or request features by opening an issue.
2. Improve documentation.
3. Submit pull requests with bug fixes or new features.

Please ensure that your code adheres to the existing style and that all tests pass before submitting a pull request.

## License

Portgen is open-source software licensed under the MIT license. See the [LICENSE](LICENSE) file for more details.
