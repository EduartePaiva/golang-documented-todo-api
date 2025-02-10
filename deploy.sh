#!/bin/zsh

set -e  # Exit on error
set -u  # Treat unset variables as an error
set -o pipefail  # Exit on first failure in a pipeline

echo "Starting build process..."

# Step 1: run local tests first
# golang tests
godotenv -f .env go test ./... -race -v
# pnpm tests
pnpm test


# Step 2: Clean previous build
echo "Cleaning previous build..."
rm -rf deploy/

# Step 3: make folders
echo "setting up folders"
mkdir -p deploy/app
# Step 4: copy static assets
echo "copying static files related to the go binary"
cp -r app/static ./deploy/app

# Step 5: build frontend app
echo "building frontend"
pnpm --dir ./frontend build
# copy dist to deploy folder
cp -r ./frontend/dist ./deploy/frontend

# Step 6: build the golang app binary

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./deploy/goapp main.go


echo "Build complete!"