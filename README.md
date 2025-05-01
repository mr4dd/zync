# Zync

Zync is a service for provisioning code running/compilation environments written in Go and utilizing Podman. This project is still under development and is not in a finalized or production-ready state.

The service supports the creation of containers for different language environments (JavaScript, Python, Go) with various runners (Node.js, Yarn, Bun, etc.).

## Features

- **Dynamic Container Creation**: Users can request containers with specific languages and runners.
- **Container Management**: Start and delete containers via API endpoints.
- **Language & Runner Support**: Support for JavaScript (Node.js, Yarn, Bun), Python, and Go.

## Installation

Clone this repository:
```
git clone https://github.com/mr4dd/zync.git
```
Navigate to the project directory:
```
cd zync
```
Install dependencies (if any):
```
go get
```
## Usage

### Start a New Container

Make a `POST` request to `/new` with the following JSON body:
```
{
  "Lang": "Javascript",
  "Runner": "Node.js"
}
```
### Delete a Container

Make a `DELETE` request to `/delete/{container_id}` to remove a specific container.

## API Endpoints

- **POST `/new`**: Starts a new container based on the specified language and runner.
- **DELETE `/delete/{id}`**: Deletes the container with the specified ID.

## Contributing

Feel free to submit issues or pull requests if you have any suggestions or improvements.

