# Hello Universe

A simple command-line application written in Go that prints a greeting based on a user-provided language code.

This project was built to practice Go fundamentals, including functions, maps, user input, string manipulation, and table-driven unit tests.

## Supported Languages

| Language | Code |
|-----------|------|
| English | `en` |
| French | `fr` |
| German | `de` |
| Indonesian | `id` |
| Italian | `it` |
| Swedish | `sv` |

If an unsupported language code is entered, the application displays an appropriate message.

## How to Run

Clone the repository and navigate to the project directory:

```bash
cd 01-hello-universe
```

Run the application:

```bash
go run .
```

Example:

```text
Enter language code: fr
Bonjour, Univers!
```

## Run Tests

```bash
go test -v
```

## Example Output

```text
Enter language code: en
Hello, Universe!

Enter language code: de
Hallo, Universum!

Enter language code: xyz
I don't know the language code: xyz
```

## Concepts Practiced

- Functions
- Maps
- User Input
- String Manipulation
- Error Handling
- Table-Driven Tests
- Go Testing Package