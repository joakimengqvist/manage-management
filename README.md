# manage-management
A hobby project for learning - An application for project and employee management.

Warning for developers who loves to reuse code, this application has a very repetitive structure. Components and functionality is only made modular if there is a clear beneficial purpose for it, and does not cause any parts of the application being less readable. The sole reason behind this is to keep the codebase as simple as possible.

## prerequisites
* **Docker:** Install Docker for your OS from the official website.
* **GNU Make:** Install via your OS package manager or a relevant tool like Homebrew (macOS) or MinGW (Windows).
* **Go (Golang):** Download and install Go from the official website.
* **Yarn:** Install Yarn via npm (Node.js package manager) after installing Node.js.

## frontend
 React app running on Vite

run yarn install in the front-end folder, and then simply run `make start` in project folder.
 
## backend
Go micro services

In the make file inside the projects folder you can find the scripts to run the go micro services

### services
* Broker-service (API service)

* Authentication-service
* Project-service
* External-companies-service
* Logger-service
* Mail-service
* Economics-service
* Notes-service
* Listener-service


