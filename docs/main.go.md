# ESH Package Execution Documentation

## Overview

The ESH (Executable Shell Script) package execution tool is a Go program designed to parse and execute ESH files. An ESH file contains a package declaration and one or more function declarations. Each function consists of a series of shell commands. The tool can execute a specific function or all functions in the package.

## Function Descriptions

### ParseEsh

*   **Description:** Parses an ESH file and returns a `Package` struct containing the package name and a slice of `Function` structs.
*   **Parameters:**
    *   `filePath`: The path to the ESH file to parse.
*   **Return values:**
    *   A `*Package` pointer if parsing is successful.
    *   An `error` if parsing fails.

### ExecuteCommand

*   **Description:** Executes a shell command and logs the output to a file.
*   **Parameters:**
    *   `cmd`: A `Command` struct containing the command name and arguments.
    *   `outputFile`: A file pointer to write the command output to.
    *   `abortChan`: A channel to listen for abort signals.
*   **Return values:**
    *   An `error` if command execution fails.

### ExecuteFunction

*   **Description:** Executes a function by running each command in its body.
*   **Parameters:**
    *   `fn`: A `*Function` pointer to the function to execute.
    *   `outputFile`: A file pointer to write the function output to.
    *   `abortChan`: A channel to listen for abort signals.
*   **Return values:**
    *   An `error` if function execution fails.

### ExecutePackage

*   **Description:** Executes a package by running all functions or a specific function.
*   **Parameters:**
    *   `pkg`: A `*Package` pointer to the package to execute.
    *   `outputFile`: A file pointer to write the package output to.
    *   `functionToRun`: The name of the function to run (optional).

## Usage Examples

### Running the ESH Tool

To run the ESH tool, use the following command:

```bash
esh <file.esh> [functionToRun]
```

Replace `<file.esh>` with the path to your ESH file. If you want to run a specific function, provide its name as the second argument.

### Creating an ESH File

Here's an example ESH file:

```esh
package mypackage

func myfunction {
  echo "Hello, World!"
  ls -l
}

func anotherfunction {
  echo "This is another function"
}
```

You can run this ESH file using the command:

```bash
esh mypackage.esh
```

This will execute both functions in the package. To run only one function, specify its name:

```bash
esh mypackage.esh myfunction
```

## Struct Definitions

### Command

*   `Name`: The name of the command.
*   `Args`: The arguments for the command.

### Function

*   `Name`: The name of the function.
*   `Body`: A slice of `Command` structs representing the function body.
*   `Executed`: A boolean indicating whether the function has been executed.
*   `ExecutedBy`: A slice of strings representing the functions that have executed this function.

### Package

*   `Name`: The name of the package.
*   `Functions`: A slice of `Function` structs representing the package functions.