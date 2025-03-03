# buf-plugin-openapi

A [buf](https://buf.build) plugin that generates OpenAPI v3 specifications from Protocol Buffer definitions. This plugin is designed to work with gRPC-Gateway annotations to create accurate OpenAPI/Swagger documentation.

## Features

- Generates OpenAPI v3.0.0 specifications from Protocol Buffer definitions
- Supports gRPC-Gateway HTTP annotations
- Converts Protocol Buffer types to OpenAPI types
- Handles path parameters and query parameters
- Supports JSON field name customization
- Generates proper request/response schemas

## Installation

```bash
go install github.com/seanmcgary/buf-plugin-openapi@latest
```

## Usage

Add the plugin to your `buf.yaml`:

```yaml
version: v1
deps:
  - buf.build/googleapis/googleapis
  - buf.build/grpc-ecosystem/grpc-gateway
plugins:
  - plugin: buf-plugin-openapi
    out: gen/openapi
    opt:
      - emit_defaults=true
      - json_names_for_fields=true
```

Available options:

- `emit_defaults`: Include default values in the OpenAPI output (default: true)
- `omit_enum_default_value`: Omit default values for enum types (default: false)
- `enums_as_ints`: Render enum values as integers instead of strings (default: false)
- `simple_operation_ids`: Remove the service prefix from operation IDs (default: false)
- `json_names_for_fields`: Use JSON names for fields instead of proto names (default: true)

## Example

Given a Protocol Buffer definition with gRPC-Gateway annotations:

```protobuf
syntax = "proto3";

package example.v1;

import "google/api/annotations.proto";

service ExampleService {
  rpc GetExample(GetExampleRequest) returns (GetExampleResponse) {
    option (google.api.http) = {
      get: "/v1/examples/{id}"
    };
  }
}

message GetExampleRequest {
  string id = 1;
}

message GetExampleResponse {
  string id = 1;
  string name = 2;
  int32 count = 3;
}
```

The plugin will generate an OpenAPI v3 specification with proper paths, parameters, and schemas.

## License

MIT 