# PortHawk

PortHawk is a lightweight, fast, and concurrent port scanning tool written in Go. It is designed to help users scan open ports on an IP address or DNS name, with options to specify a port range and limit the number of concurrent scanning routines.

---

## Features

- **Fast concurrent port scanning**: Utilizes Go routines to scan ports simultaneously.
- **Controlled concurrency**: Limit the number of parallel routines with a semaphore for stable performance.
- **Customizable port range**: Specify the start and end ports for scanning.
- **Cross-platform**: Runs on any system that supports Go.

---

## Installation

### **Prerequisites**
- [Go](https://go.dev/) (version 1.23 or later) installed on your system.
- Git installed to clone the repository.

### **Clone the repository**
```bash
git clone https://github.com/Krikas-Sec/PortHawk.git
cd PortHawk
```
Build the program
Navigate to the project directory:
```bash
cd PortHawk
```
Build an executable file:
```bash
go build -o porthawk
```
Test the program:
```bash
./porthawk
Usage
Run the program
Run the program with the following command:
```
```bash
./porthawk
```
Example usage
Enter the IP address or DNS name:
```bash
Enter IP address or hostname to scan: 127.0.0.1
Specify the port range:
```
```bash
Enter the start port: 20
Enter the end port: 100
```
Limit the number of concurrent routines:
```bash
Enter the maximum number of concurrent routines (e.g., 100): 10
```
The program will scan the ports and list all open ports.

Example Output
```bash
Welcome to PortHawk - Enhanced Port Scanner
Enter IP address or hostname to scan: 127.0.0.1
Enter the start port: 20
Enter the end port: 100
Enter the maximum number of concurrent routines (e.g., 100): 10
Scanning ports 20 to 100 on 127.0.0.1...
Port 22 is open
Port 80 is open
Scan complete.
```
Technical Overview
Architecture
The program uses Go routines for concurrent port scanning.
A semaphore (implemented with channels) limits the number of simultaneous routines to optimize resource usage.
WaitGroup ensures the program waits for all routines to complete before proceeding.
File Structure
```bash
PortHawk/
├── main.go        # Main file for the program
├── go.mod         # Go module files
```
Future Enhancements
Log scan results to a file.
Add support for UDP port scanning.
Create a configuration file for default settings.
Contributing
If you would like to contribute to this project:

Fork the repository.
Create a new branch:
```bash
git checkout -b feature-name
```
Make your changes and commit:
```bash
git commit -m "Added feature-name"
```
Submit a pull request.
License
This project is licensed under the MIT License. See LICENSE for more information.

Support
If you appreciate this tool and want to support its development, please visit my Buy Me a Coffee page https://buymeacoffee.com/Tempcoder. Thank you for your support!


