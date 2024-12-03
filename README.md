# Advent of Code 2024

This repository contains solutions to the [Advent of Code 2024](https://adventofcode.com/2024) problems, implemented in **Go**. Each day's solution is organized in its own folder with the corresponding input files and Docker configurations for easy execution.

## Project Structure

The repository is structured as follows:

```
.
|-- dayX/
|   |-- adventCode1.go      # Solution for Day X
|   |-- Dockerfile          # Docker configuration for Day X
|   |-- input1.txt          # Input file for Day X
|
|-- README.md               # Project documentation
```

Each folder contains:
- **`adventCodeX.go`**: The Go solution for the specific day's problem.
- **`inputX.txt`**: The input file for the problem.
- **`Dockerfile`**: Docker configuration to run the solution in a containerized environment.

---

## Prerequisites

### Without Docker:
- Go installed (at least version 1.20).
- A terminal or command-line environment.

### With Docker:
- Docker installed and running.

---

## How to Execute

### Without Docker
1. Navigate to the specific day's folder:
   ```bash
   cd dayX
   ```
2. Run the Go program:
   ```bash
   go run adventCodeX.go
   ```
   
   Ensure the corresponding `inputX.txt` file is in the same directory as the `.go` file.

### With Docker
1. Navigate to the specific day's folder:
   ```bash
   cd dayX
   ```
2. Build the Docker image:
   ```bash
   docker build -t adventcodeX .
   ```
3. Run the Docker container:
   ```bash
   docker run --rm adventcodeX
   ```

---

## Notes
- The solutions are designed to read from the input files (e.g., `inputX.txt`) located in the same directory.
- Ensure your input files are correctly formatted as specified in the Advent of Code problems.

Happy Coding! 🎄

