* goapp-boilerplate

boilerplate go api stuff

#### References

generates types and handlers from api spec:

oapi-codegen README: https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file

TLDR;
Define your spec in the spec yaml files, then run `make gen` to generate the types and handlers.

generates docs from api spec:

swagger UI: https://github.com/flowchartsman/swaggerui

#### Running locally


IMPORTANT!
update the `go:generate` tag in the api implementation file `impl.go` with your file paths prior to running `go generate` on your machine

run `make dev` to build and run the container locally. By default it will use port 3000. If you change the configuration,
be sure to update the dockerfile and docker compose with any port changes.

#### Testing

run `make test` to run tests, or run them manually with the `go test` command
