# OSCParser

**OSCParser** is a Go package designed for efficient parsing and extraction of variable names from `.osc` files. It provides a straightforward interface to retrieve all variables, making it an essential tool for handling and processing `.osc` data in your Go applications.

## Features
- Parse `.osc` files to extract variable names
- Returns a map of variable names found in the file
- Easy to integrate and use in your Go projects

## Installation

To use the `oscparser` package in your project, you need to import it directly from your GitHub repository. First, ensure your project is using Go modules (`go.mod`).

1. Initialize your Go project (if not already done):
  ```bash
    go mod init exampleProject
  ```

2. Import the package in your project:
  ```Go
    import (
      "github.com/f0xb17/OSCParser"
    )
  ```

## Usage

  ```Go
    package main

    import (
      "fmt"
      "github.com/f0xb17/OSCParser"
    )

    func main() {
      variables, err := oscparser.Read_File("path/to/your/file", `\((S\.L|L\.L)\.(\w+)\)`)
      if err != nil {
        fmt.Println(err)
      }

      for variable := range variables {
        fmt.Println(variable)
      }
    }
  ```
## Contributing
Feel free to submit issues or pull requests if you find any bugs or want to add new features. Contributions are always welcome!
