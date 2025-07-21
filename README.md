# 🚀 goapp-boilerplate

A modern Go API boilerplate with embedded React UI, featuring OpenAPI specification, Swagger documentation, and Prometheus metrics.

## ✨ Features

- **🏗️ Go Backend**: Built with modern Go patterns and Cobra CLI
- **⚛️ React Frontend**: Statically built and embedded into the Go binary
- **📋 OpenAPI Specification**: API-first development with oapi-codegen
- **📖 Swagger UI**: Interactive API documentation
- **📊 Prometheus Metrics**: Built-in monitoring and metrics
- **🐳 Docker Support**: Multi-stage builds with UI compilation
- **🔧 Hot Reload**: Development scripts for rapid iteration

## 🚀 Quick Start

### Prerequisites

- **Go** 1.22 or later
- **Node.js** 18 or later
- **Docker** (optional, for containerized development)

### Local Development

1. **Quick Start (Recommended)**:
   ```bash
   ./dev.sh
   ```

2. **Manual Build**:
   ```bash
   # Build UI and Go application
   make build-local
   
   # Run the application
   ./bin/goapp-boilerplate serve
   ```

3. **Using Docker**:
   ```bash
   make dev
   ```

### 🌐 Access Points

Once running, your application will be available at:

- **🏠 Main UI**: http://localhost:3000
- **📖 API Documentation**: http://localhost:3000/swagger/
- **📊 Metrics**: http://localhost:3000/metrics
- **🔍 API Endpoint**: http://localhost:3000/ping

## 🏗️ Project Structure

```
├── api/                    # OpenAPI specification and generated code
│   ├── api.yml            # OpenAPI 3.0 specification
│   ├── config.yml         # oapi-codegen configuration
│   ├── api.gen.go         # Generated API code
│   └── impl.go            # API implementation
├── cmd/                   # CLI commands (Cobra)
├── config/                # Configuration management
├── server/                # HTTP server setup
├── static/                # Embedded static files (auto-generated)
├── ui/                    # React frontend source
│   ├── src/               # React components and logic
│   ├── public/            # Public assets
│   └── package.json       # Node.js dependencies
├── Dockerfile             # Multi-stage Docker build
├── docker-compose.yml     # Local development environment
├── Makefile              # Build automation
└── dev.sh                # Development script
```

## 🛠️ Development

### Building Components

```bash
# Build only the UI
make build-ui

# Build everything locally
make build-local

# Build Docker image
make build

# Clean all build artifacts
make clean
```

### Running Tests

```bash
make test
```

### Adding New API Endpoints

1. **Update the OpenAPI specification** in `api/api.yml`
2. **Regenerate the code**:
   ```bash
   go generate ./api
   ```
3. **Implement the handler** in `api/impl.go`
4. **Test your endpoint** using the built-in UI or Swagger documentation

## 📋 API Development

This project uses **API-first development** with OpenAPI 3.0:

### Code Generation

The project uses [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) to generate types and handlers from the API specification.

**Important**: Update the `go:generate` tag in `api/impl.go` with your correct file paths before running `go generate`.

### Workflow

1. Define your API in `api/api.yml`
2. Run `go generate ./api` to generate types and handlers
3. Implement handlers in `api/impl.go`
4. Test using the React UI or Swagger documentation

## 🐳 Docker

### Local Development

```bash
# Start with Docker Compose
make dev

# Or manually
docker-compose up go-api
```

### Production Build

The Dockerfile uses multi-stage builds:

1. **UI Build Stage**: Compiles React app with Node.js
2. **Go Build Stage**: Compiles Go binary with embedded UI
3. **Runtime Stage**: Minimal Ubuntu image with the final binary

## 🎨 UI Customization

The React UI is located in the `ui/` directory. It's a standard Create React App with:

- **Real-time API testing** interface
- **Responsive design** with modern CSS
- **Direct links** to Swagger UI and metrics
- **Development guidelines** and next steps

To customize:

1. Modify files in `ui/src/`
2. Run `npm start` in the `ui/` directory for hot reload
3. Build with `make build-ui` when ready

## 📊 Monitoring

Built-in Prometheus metrics are available at `/metrics`, including:

- HTTP request metrics
- Response time histograms
- Error rates
- Custom business metrics (add your own!)

## 🔧 Configuration

Configuration is managed through YAML files in the `config/` directory:

- `defaultConfig.yaml`: Default settings
- Override with `--config` flag or environment variables

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass with `make test`
5. Submit a pull request

## 📝 References

- **oapi-codegen**: https://github.com/oapi-codegen/oapi-codegen
- **Swagger UI**: https://github.com/flowchartsman/swaggerui
- **OpenAPI 3.0**: https://swagger.io/specification/
- **Go Embed**: https://pkg.go.dev/embed

## 📄 License

This project is licensed under the terms found in the LICENSE file.
