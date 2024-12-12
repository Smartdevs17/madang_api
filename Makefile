# Makefile for madang_api

# Set the environment variable
export GO_ENV=development

# Default target
run:
	@echo "Running madang_api in development mode..."
	@if lsof -i :3000; then \
		echo "Port 3000 is already in use. Please free the port and try again."; \
		exit 1; \
	fi
	export PATH=$PATH:$(go env GOPATH)/bin
	# source ~/.zshrc
	CompileDaemon -command="./madang_api" -verbose

# Clean target (optional, if you have any build artifacts to clean)
clean:
	@echo "Cleaning up..."
	# Add any cleanup commands here

.PHONY: run clean
