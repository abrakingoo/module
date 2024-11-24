# Lem-in

A Go implementation of an ant colony pathfinding simulator.

## Overview

Lem-in is a program that simulates ants navigating through an ant colony to find the most efficient path from start to end. The program reads a colony description from a file and outputs the optimal sequence of ant movements.

## Features

- Efficient pathfinding using Breadth-First Search (BFS)
- Handles multiple valid paths
- Input validation and error handling
- Optimized for minimal moves
- Detailed movement visualization

## Requirements

- Go 1.19 or higher
- Standard Go packages only

## Installation

```bash
git clone https://learn.zone01kisumu.ke/git/abrakingoo/lem-in
cd lem-in
```

## Usage

Run the program with an input file:

```bash
go run ./cmd [filename]
```

### Input File Format

The input file should contain:
1. Number of ants
2. Room definitions (name + coordinates)
3. Tunnel definitions (connections between rooms)

Example input:
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

### Output Format

The program outputs:
1. Input file content
2. Sequence of ant movements in format `Lx-y` where:
   - x: ant number
   - y: destination room name

Example output:
```
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```

## Project Structure

- `/data`: Core data structures and parsing
- `/utils`: Utility functions
- `/tests`: Test files and test data

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test ./... -cover
```

## Rules and Constraints

1. Room Rules:
   - No names starting with 'L' or '#'
   - No spaces in room names
   - Integer coordinates only

2. Tunnel Rules:
   - Connect exactly two rooms
   - No duplicate tunnels between same rooms
   - Each tunnel usable once per turn

3. Movement Rules:
   - One ant per room (except start/end)
   - Each ant moves once per turn
   - No ant collisions

## Error Handling

The program handles various error cases:
- Invalid number of ants
- Missing start/end rooms
- Invalid room formats
- Invalid connections
- Unreachable end room

Error messages are descriptive and specific to the issue encountered.

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Project inspired by ant colony optimization algorithms
- Thanks to all contributors and testers
