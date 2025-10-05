# Flint

A simple and non-intrusive tool to deploy and manage your containerized applications with Docker on remote servers.

## Philosophy

Flint was designed with a core idea: simplicity and respect for your server environment. Unlike other deployment tools that take full control of your machine, Flint is designed to be as unintrusive as possible.

### Key Points:

*   **✅ No Root Access:** No need for `sudo` or `root` permissions to operate. The tool runs entirely in user space.
*   **🐋 Docker-Based:** Leverages the power and isolation of Docker to deploy your applications.
*   **🚀 Automated Reverse Proxy:** A reverse proxy is automatically configured to expose your services, without complex manual configuration.
*   **📂 No Database:** No database to install or maintain. Configuration is managed through simple files.
*   **💡 Lightweight and Simple:** A single application to run to manage the entire lifecycle of your deployments.

## How It Works

Flint is an application written in **Go** that exposes a REST API (via the **Gin** framework) to manage deployments. It works by connecting to your remote server (e.g., via SSH) and using the machine's Docker API to manage containers.

1.  **Configuration:** You describe your applications in a simple configuration file (e.g., `docker-compose.yml` or a proprietary format).
2.  **Deployment:** The tool transfers the configuration to the remote server.
3.  **Launch:** It starts your Docker containers based on this configuration.
4.  **Routing:** It dynamically updates the reverse proxy's configuration (itself a Docker container) to route traffic to your new applications.

All of this without ever requiring elevated privileges on the host machine.

## Installation

*(Installation instructions to come)*

## Quick Start

### Running the server locally

To run the Flint server on your local machine, execute the following command:

```bash
go run main.go
```

### Deployment Example

*(Usage example to come)*

```bash
# Example command (to be defined)
flint deploy --file my-app.yml --target user@my-server.com
```

## Contributing

Contributions are welcome! Feel free to open an issue or a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.