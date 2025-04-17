# Development Environment Setup

## Prerequisites

- Docker
- Docker Compose

## Configuration

1. Create a `.env` file in the docker/dev/ folder and the docker/prod/ folder. This file will contain environment variables used by Docker Compose and the application.

   ```
   arly/
   ├── .github/workflows/
   │   └── ci.dev.backend.yml
   ├── docker/
   │   ├── dev/
   │   │   ├── .env
   │   │   └── docker-compose.yml
   │   └── prod/
   │       ├── .env
   │       └── docker-compose.yml
   ├── backend/
   │   └── ...
   ├── frontend/
   │   └── ...
   ├── scripts/
   │   └── dev.launch-compose.sh
   └── ...
   ```

2. Add the necessary environment variables to the `.env` file. For example:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=myapp
   DB_USER=myuser
   DB_PASSWORD=mypassword
   ...
   ```

## Launching Docker Compose in Development Mode

We provide scripts in the `scripts/` folder to simplify the process of launching Docker Compose in development mode.

1. Open a terminal and navigate to the project root directory.

2. Run the development script:

   ```
   ./scripts/dev.launch-compose.sh
   ```

   This script will:
   - Source the `.env` file
   - Build the necessary Docker images (if needed)
   - Start the Docker Compose services in development mode
   - Display the logs from all services

3. To stop the services, press `Ctrl+C` in the terminal where the script is running.

4. To completely tear down the development environment, including volumes, run:

   ```
   ./scripts/dev-down.sh
   ```
## Troubleshooting

If you encounter any issues:

1. Ensure your `.env` file is correctly configured
2. Check the Docker and Docker Compose logs for any error messages
3. Verify that all required ports are available on your machine

For further assistance, please contact Illyas CHIHI on Teams.
