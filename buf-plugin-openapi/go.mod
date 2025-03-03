module github.com/Layr-Labs/buf-plugin-openapi

go 1.23.6

require (
	github.com/Layr-Labs/protocol-apis-annotations v0.0.0-00010101000000-000000000000
	google.golang.org/genproto/googleapis/api v0.0.0-20240123012728-ef4313101c80
	google.golang.org/protobuf v1.36.5
)

require google.golang.org/genproto v0.0.0-20240123012728-ef4313101c80 // indirect

replace github.com/Layr-Labs/protocol-apis-annotations => ../protocol-apis-annotations
