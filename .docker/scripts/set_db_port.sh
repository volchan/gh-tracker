#!/usr/bin/env bash

# This script updates the DB_PORT in the .env file if the docker container port has changed
APP_ENV_FILE_PATH="../.env"
DOCKER_ENV_FILE_PATH=".env"

# Get the new port from the docker compose command
NEW_PORT=$(docker compose port db 5432 | cut -d: -f2)

# Create .env file if it doesn't exist
[ -f "$APP_ENV_FILE_PATH" ] || task -t ../Taskfile.yml env:setup

# Get the current port from the .env file
CURRENT_PORT=$(awk -F= '/^DB_PORT=/{print $2}' "$APP_ENV_FILE_PATH")

# Check if the port has changed
if [ "$CURRENT_PORT" = "$NEW_PORT" ]; then
    echo "Database still running on port $CURRENT_PORT. No changes needed."
    exit 0
fi

# Update the DB_PORT in the .env file
sed -i.bak "s/^DB_PORT=.*/DB_PORT=$NEW_PORT/" "$APP_ENV_FILE_PATH" && rm "${APP_ENV_FILE_PATH}.bak"
sed -i.bak "s/^POSTGRES_PORT=.*/POSTGRES_PORT=$NEW_PORT/" "$DOCKER_ENV_FILE_PATH" && rm "${DOCKER_ENV_FILE_PATH="../.env"}.bak"
echo "DB_PORT has been updated to $NEW_PORT in the .env file"
