FROM node:18-alpine AS ui-builder

WORKDIR /ui

# Copy package.json and package-lock.json (if available)
COPY ui/package*.json ./

# Install dependencies
RUN npm ci --only=production

# Copy UI source code
COPY ui/ ./

# Build the React app
RUN npm run build

FROM golang:latest AS base

WORKDIR /svc

ARG GIT_CREDS
ENV GIT_CREDS=${GIT_CREDS}
RUN set -eux; \
  git config --global credential.helper store; \
  git config --global --unset-all 'url.ssh://git@github.com.insteadOf' || echo "No ssh config"; \
  echo "https://${GIT_CREDS}@github.com" > ~/.git-credentials

COPY go.mod go.sum /svc/
RUN go mod download

# Copy the built UI into the static directory
COPY --from=ui-builder /ui/build /svc/static/ui

COPY . /svc
RUN CGO_ENABLED=0 GOOS=linux go build -o /svc/dist/app

FROM ubuntu:20.04 AS app

EXPOSE 3000
COPY --from=base /svc/dist/app /app
COPY --from=base /svc/config/defaultConfig.yaml /config/defaultConfig.yaml
CMD ["/app", "serve"]

