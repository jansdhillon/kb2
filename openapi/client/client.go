package client

//go:generate sh -c "set -e; if [ -n \"$OPENAPI_SPEC\" ]; then if [ ! -f \"$OPENAPI_SPEC\" ]; then echo \"missing OpenAPI spec: $OPENAPI_SPEC\" >&2; exit 1; fi; exec go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml \"$OPENAPI_SPEC\"; else if [ ! -f ../openapi.bundle.yaml ]; then echo \"missing OpenAPI spec: ../openapi.bundle.yaml\" >&2; exit 1; fi; exec go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../openapi.bundle.yaml; fi"
