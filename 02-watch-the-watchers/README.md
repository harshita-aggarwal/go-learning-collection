# Watch the Watchers

A Go application for managing and analyzing movie watchers stored in JSON files.

This project was built to practice working with JSON data, file I/O, custom structs, slices, maps, sorting, and unit testing in Go.

## Features

- Load watcher data from JSON files
- Find popular movies watched by multiple users
- Check whether a watcher already exists using a username or email
- Merge multiple watcher lists while preventing duplicate users
- Export watcher data back to JSON
- Comprehensive unit test coverage

## Project Structure

```text
02-watch-the-watchers/
├── cmd/
│   └── main.go
├── watcher/
│   ├── watcher.go
│   └── watcher_test.go
├── test_data/
│   ├── watchers.json
│   ├── watchers-modified.json
│   └── watchers-contains-error.json
├── go.mod
└── go.sum
```

## How to Run

Navigate to the project directory:

```bash
cd 02-watch-the-watchers
```

Run the application:

```bash
go run ./cmd
```

## Run Tests

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## Example Output

```text
[1 of 6] Username: movie.lover / Email: movie.lover@example.com
[2 of 6] Username: film.fanatic / Email: film.fanatic@example.com
[3 of 6] Username: cinema.critic / Email: cinema.critic@example.com
[4 of 6] Username: movie.always / Email: movie.always@example.com
[5 of 6] Username: new.watcher2 / Email: new.watcher2@example.com
[6 of 6] Username: new.watcher3 / Email: new.watcher3@example.com
```

## Concepts Practiced

- Structs and Custom Types
- JSON Encoding and Decoding
- File I/O
- Maps and Slices
- UUIDs
- Sorting with Custom Comparators
- String Normalization
- Package Organization
- Error Handling
- Table-Driven Tests
- Unit Testing

## Key Learnings

This project helped reinforce how to model real-world data using Go structs, work with JSON files, organize code into packages, write reusable utility functions, and build reliable test suites using table-driven testing patterns.