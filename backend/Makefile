# Define variables
GO_VERSION ?= 1.23.4
GO_BINARY := go
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')

# Define log file
LOG_FILE := logs/makefile.log

# Function to log messages
log:
	@mkdir -p logs
	@touch $(LOG_FILE)
	@echo "[$(shell date)] $1" >> $(LOG_FILE)

# Print a message with a timestamp
timestamped_echo:
	@echo "[$(shell date)] $1"

check-env:
	@if [ ! -f .env ]; then \
		echo "No .env file found. Please create a .env file with the required environment variables." | tee -a $(LOG_FILE); \
		exit 1; \
	else \
		echo ".env file found." | tee -a $(LOG_FILE); \
	fi

# Check if Go is installed
check-go:
	@echo "Checking if Go is installed..."
	@if which $(GO_BINARY) > /dev/null 2>&1; then \
		echo "Go is already installed." | tee -a $(LOG_FILE); \
	else \
		echo "Go is not installed. Please install Go $(GO_VERSION)" | tee -a $(LOG_FILE); \
		exit 1; \
	fi
# Install Go (for Linux and macOS)
install-go:
	@echo "Installing Go $(GO_VERSION)..."
	@if [ "$(OS)" = "linux" ]; then \
		curl -OL https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz && \
		sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz && \
		if ! grep -q 'export PATH=$$PATH:/usr/local/go/bin' ~/.bashrc; then \
			echo 'export PATH=$$PATH:/usr/local/go/bin' >> ~/.bashrc; \
		fi && \
		source ~/.bashrc && \
		echo "Go $(GO_VERSION) installed successfully on Linux." | tee -a $(LOG_FILE); \
	elif [ "$(OS)" = "darwin" ]; then \
		curl -OL https://golang.org/dl/go$(GO_VERSION).darwin-amd64.pkg && \
		sudo installer -pkg go$(GO_VERSION).darwin-amd64.pkg -target / && \
		if ! grep -q 'export PATH=$$PATH:/usr/local/go/bin' ~/.zshrc; then \
			echo 'export PATH=$$PATH:/usr/local/go/bin' >> ~/.zshrc; \
		fi && \
		source ~/.zshrc && \
		echo "Go $(GO_VERSION) installed successfully on macOS." | tee -a $(LOG_FILE); \
	else \
		echo "OS $(OS) is not supported for automatic Go installation" | tee -a $(LOG_FILE); \
		exit 1; \
	fi

# Check if swag is installed
check-swag:
	@echo "Checking if swag is installed..."
	@if which swag > /dev/null 2>&1; then \
		echo "Swag is already installed." | tee -a $(LOG_FILE); \
	else \
		echo "Swag is not installed. Installing swag..." | tee -a $(LOG_FILE); \
		make install-swag; \
	fi
# Install swag
install-swag:
	@echo "Installing swag..."
	@if [ -z "$(GO_BINARY)" ]; then \
		echo "GO_BINARY is not set. Please set it to the path of the Go binary (e.g., 'go')." | tee -a $(LOG_FILE); \
		exit 1; \
	fi
	@$(GO_BINARY) install github.com/swaggo/swag/cmd/swag@latest
	@if ! grep -q 'export PATH=$$PATH:$$HOME/go/bin' ~/.bashrc; then \
		echo 'export PATH=$$PATH:$$HOME/go/bin' >> ~/.bashrc; \
	fi
	@source ~/.bashrc
	@which swag > /dev/null 2>&1 && echo "Swag installed successfully." | tee -a $(LOG_FILE) || (echo "Failed to install swag." | tee -a $(LOG_FILE) && exit 1)

# Check if air is installed
check-air:
	@echo "Checking if air is installed..."
	@if which air > /dev/null 2>&1; then \
		echo "Air is already installed." | tee -a $(LOG_FILE); \
	else \
		echo "Air is not installed. Installing air..." | tee -a $(LOG_FILE); \
		make install-air; \
	fi

# Install air
install-air:
	@echo "Installing air..."
	@$(GO_BINARY) install github.com/air-verse/air@latest
	@which air > /dev/null 2>&1 && echo "Air installed successfully." | tee -a $(LOG_FILE) || (echo "Failed to install air." | tee -a $(LOG_FILE) && exit 1)

# Install dependencies
install-deps:
	@echo "Installing project dependencies..."
	@$(GO_BINARY) mod tidy
	@echo "Project dependencies installed." | tee -a $(LOG_FILE)

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init
	@echo "Swagger documentation generated." | tee -a $(LOG_FILE)

# Targets
build-windows:
	@echo "Building the project for Windows..."
	@$(GO_BINARY) build -o ./build/arlyApi-windows.exe main.go
	@echo "Project built successfully for Windows." | tee -a $(LOG_FILE)

build-linux:
	@echo "Building the project for Linux..."
	@GOOS=linux GOARCH=amd64 $(GO_BINARY) build -o ./build/arlyApi-linux main.go
	@echo "Project built successfully for Linux." | tee -a $(LOG_FILE)
	@chmod +x ./build/arlyApi-linux

build-macos:
	@echo "Building the project for macOS..."
	@GOOS=darwin GOARCH=amd64 $(GO_BINARY) build -o ./build/arlyApi-macos main.go
	@echo "Project built successfully for macOS." | tee -a $(LOG_FILE)
	@chmod +x ./build/arlyApi-macos

check-env:
	@if [ -f "./build/.env" ]; then \
		echo "Checking if .env file has changed..."; \
		if cmp -s .env ./build/.env; then \
			echo ".env file is up to date."; \
		else \
			echo ".env file has changed. Copying new .env to build directory..."; \
			cp .env ./build/.env; \
			echo ".env file updated in build directory."; \
		fi \
	else \
		echo "No .env file found in build directory. Copying..."; \
		cp .env ./build/.env; \
		echo ".env file copied to build directory."; \
	fi

build:
	@echo "Creating the build directory..."
	@mkdir -p build
	@echo "Build directory created." | tee -a $(LOG_FILE)
	@$(MAKE) check-env
	@echo "Building the project for all platforms..."
	@$(MAKE) build-windows
	@$(MAKE) build-linux
	@$(MAKE) build-macos
	@echo "All platform builds completed successfully." | tee -a $(LOG_FILE)

# Run the project
run:
	@echo "Running the project..."
	@$(GO_BINARY) run main.go
	@echo "Project is running." | tee -a $(LOG_FILE)

# Run the project with air
dev: check-air
	@echo "Running the project with air..."
	@air
	@echo "Project is running with air for live reloading." | tee -a $(LOG_FILE)

# Clean up
clean:
	@echo "Cleaning up build artifacts..."
	@rm -f myapp
	@echo "Clean up completed." | tee -a $(LOG_FILE)

# Check Go installation and install dependencies
setup: check-go install-deps check-swag check-air check-env

# Default target
all: setup build swagger

.PHONY: check-go install-go install-deps check-swag install-swag check-air install-air swagger build run dev clean setup all build-windows build-linux build-macos 
