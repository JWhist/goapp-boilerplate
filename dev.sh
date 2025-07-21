#!/bin/bash

# Development script for the Go API with React UI

set -e

echo "🚀 Starting Go API Boilerplate with UI"

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js to build the UI."
    exit 1
fi

# Check if npm is installed
if ! command -v npm &> /dev/null; then
    echo "❌ npm is not installed. Please install npm to build the UI."
    exit 1
fi

# Build the UI
echo "📦 Building React UI..."
cd ui
if [ ! -d "node_modules" ]; then
    echo "📥 Installing UI dependencies..."
    npm install
fi
npm run build
cd ..

# Copy UI build to static directory
echo "📁 Copying UI build to static directory..."
rm -rf static/ui
mkdir -p static
cp -r ui/build static/ui

# Build and run the Go application
echo "🔨 Building Go application..."
go build -o bin/goapp-boilerplate ./main.go

echo "🎉 Build complete!"
echo "🌐 Starting server..."
echo "📱 UI will be available at: http://localhost:3000"
echo "📖 API documentation at: http://localhost:3000/swagger/"
echo "📊 Metrics at: http://localhost:3000/metrics"
echo ""

./bin/goapp-boilerplate serve
