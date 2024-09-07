# How to use this app in all platforms

1. Clone the repository to your local machine.
2. Navigate to the directory where the repository is cloned.
3. Run the following command to build the application:
   - On Windows: `go build -o esh.exe`
   - On Linux/Mac: `go build -o esh`
4. Once the build is successful, you can run the application with the following command:
   - On Windows: `esh.exe <file.esh>`
   - On Linux/Mac: `./esh <file.esh>`
5. The application will execute the commands in the .esh file and write the output to a file named `output.log` in the current directory.
6. You can check the execution status and output in the `output.log` file.

7. **Cross-Platform Compatibility**: Our `.esh` files seamlessly support both PowerShell commands on Windows and shell commands on Linux/Mac, ensuring a consistent and reliable execution experience across different operating systems.

8. **Development Status**: Our application is currently in development. We appreciate any contributions or feedback that can help us improve and expand its capabilities.

**Example Usage on Windows**

Suppose we have a file named `example.esh` with the following content:
```
package example

func build {
    powershell "Write-Host 'Building on Windows'"
    powershell "Write-Host 'Compiling code for Windows'"
    powershell "Write-Host 'Linking libraries for Windows'"
    powershell "Write-Host 'Creating executable for Windows'"
}

func test {
    powershell "Write-Host 'Running tests on Windows'"
    powershell "Write-Host 'Unit tests for Windows'"
    powershell "Write-Host 'Integration tests for Windows'"
    powershell "Write-Host 'System tests for Windows'"
}

func deploy {
    powershell "Write-Host 'Deploying on Windows'"
    powershell "Write-Host 'Preparing deployment package for Windows'"
    powershell "Write-Host 'Uploading to server for Windows'"
    powershell "Write-Host 'Configuring environment for Windows'"
}
```
To execute this file on Windows, follow these steps:

1. Save the `example.esh` file in a directory of your choice.
2. Open a Command Prompt or PowerShell in the directory where the file is saved.
3. Run the following command to execute the file: `esh example.esh`
4. The application will execute the commands in the `example.esh` file and write the output to a file named `output.log` in the current directory.

**Example Usage on Mac/Linux**

Suppose we have a file named `example.esh` with the following content:
```
package example

func build {
    bash "echo 'Building on Mac/Linux'"
    bash "echo 'Compiling code for Mac/Linux'"
    bash "echo 'Linking libraries for Mac/Linux'"
    bash "echo 'Creating executable for Mac/Linux'"
}

func test {
    bash "echo 'Running tests on Mac/Linux'"
    bash "echo 'Unit tests for Mac/Linux'"
    bash "echo 'Integration tests for Mac/Linux'"
    bash "echo 'System tests for Mac/Linux'"
}

func deploy {
    bash "echo 'Deploying on Mac/Linux'"
    bash "echo 'Preparing deployment package for Mac/Linux'"
    bash "echo 'Uploading to server for Mac/Linux'"
    bash "echo 'Configuring environment for Mac/Linux'"
}
```
To execute this file on Mac/Linux, follow these steps:

1. Save the `example.esh` file in a directory of your choice.
2. Open a Terminal in the directory where the file is saved.
3. Run the following command to execute the file: `./esh example.esh`
4. The application will execute the commands in the `example.esh` file and write the output to a file named `output.log` in the current directory.









