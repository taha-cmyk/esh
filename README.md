## Introducing `esh` - A Cross-Platform Scripting Tool


`esh` is a lightweight, flexible scripting tool designed for task automation across platforms. It enables developers to define functions in a `.esh` file and execute them in their development environment, similar to how Makefiles work. `esh` is versatile, allowing you to create reusable scripts that work seamlessly on both Windows and Unix-based systems (Linux/Mac).

### Key Features:
- **Cross-platform support** for both PowerShell (Windows) and shell (Linux/Mac) commands.
- **Custom script execution**: Easily automate common tasks like building, testing, and cleaning up code.
- **Modular functions**: Define reusable tasks and execute them either as a full script or individually.

A **binary release** will be available soon, making it even easier to get started without needing to build `esh` from source.

---

## How to Use `esh` on All Platforms

### 1. Clone the Repository
```bash
git clone <repository-url>
```

### 2. Build the Application
Navigate to the cloned directory and build the `esh` executable based on your operating system:

- **Windows**:
   ```bash
   go build -o esh.exe
   ```

- **Linux/Mac**:
   ```bash
   go build -o esh
   ```

### 3. Execute the Application
To run `esh`, provide a `.esh` script file as input:

- **Windows**:
   ```bash
   esh.exe <file.esh>
   ```

- **Linux/Mac**:
   ```bash
   ./esh <file.esh>
   ```

The application will execute the tasks defined in the `.esh` file and log the output to `output.log` in the current directory.

---

### Cross-Platform Compatibility

`esh` ensures that your scripts run consistently, whether you're using PowerShell commands on Windows or shell commands on Linux/Mac. It abstracts platform differences, allowing you to focus on writing the logic for your tasks.

---

### Development Status

We are continuously working to improve `esh` and add more functionality. Contributions and feedback are welcome to help shape the future of this tool.

---

### Example Usage

#### Windows Example

Here’s an example of a `.esh` script (`example.esh`) for a Windows environment:

```esh
package automation

func build {
    # Example task
    msbuild project.sln
}

func deploy {
    # Example task
    Copy-Item -Path 'build/output' -Destination 'C:/deploy'
}

func clean {
    # Example task
    Remove-Item -Force output/*
}
```

To execute the script:

1. Save `example.esh` in your working directory.
2. Open **Command Prompt** or **PowerShell** in that directory.
3. Run:
   ```bash
   esh example.esh
   ```

#### Mac/Linux Example

For Mac/Linux users, here’s an example `.esh` script:

```esh
package automation

func build {
    # Example task
    make build
}

func deploy {
    # Example task
    cp -R ./build/output /var/www/
}

func clean {
    # Example task
    rm -rf ./output/*
}
```

To execute the script:

1. Save `example.esh` in your directory.
2. Open **Terminal** in that directory.
3. Run:
   ```bash
   ./esh example.esh
   ```

---

### Running Specific Functions

With `esh`, you can execute specific functions defined in the `.esh` file, allowing for more fine-grained control over the tasks you want to run.

#### Windows Example

To execute a specific function (e.g., `build`):

1. Save `example.esh` in your directory.
2. Open **Command Prompt** or **PowerShell**.
3. Run:
   ```bash
   esh example.esh build
   ```

#### Mac/Linux Example

To run a specific function on Mac/Linux:

1. Save `example.esh` in your directory.
2. Open **Terminal** in that directory.
3. Run:
   ```bash
   ./esh example.esh build
   ```