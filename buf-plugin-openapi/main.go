package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

	openapiAnnotations "github.com/Layr-Labs/protocol-apis-annotations/protos/annotations"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

type OpenAPISpec struct {
	OpenAPI    string              `json:"openapi"`
	Info       OpenAPIInfo         `json:"info"`
	Paths      map[string]PathItem `json:"paths"`
	Components *Components         `json:"components,omitempty"`
	Tags       []Tag               `json:"tags,omitempty"`
}

type OpenAPIInfo struct {
	Title          string       `json:"title"`
	Version        string       `json:"version"`
	Description    string       `json:"description,omitempty"`
	TermsOfService string       `json:"termsOfService,omitempty"`
	Contact        *ContactInfo `json:"contact,omitempty"`
	License        *LicenseInfo `json:"license,omitempty"`
}

type PathItem struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
}

type Operation struct {
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	OperationID string              `json:"operationId"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
	RequestBody *RequestBody        `json:"requestBody,omitempty"`
	Responses   map[string]Response `json:"responses"`
	Tags        []string            `json:"tags,omitempty"`
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"` // path, query, header, cookie
	Required    bool   `json:"required"`
	Schema      Schema `json:"schema"`
	Description string `json:"description,omitempty"`
}

type RequestBody struct {
	Required bool                       `json:"required"`
	Content  map[string]MediaTypeObject `json:"content"`
}

type Response struct {
	Description string                     `json:"description"`
	Content     map[string]MediaTypeObject `json:"content,omitempty"`
}

type MediaTypeObject struct {
	Schema Schema `json:"schema"`
}

type Schema struct {
	Type       string            `json:"type,omitempty"`
	Format     string            `json:"format,omitempty"`
	Ref        string            `json:"$ref,omitempty"`
	Properties map[string]Schema `json:"properties,omitempty"`
	Items      *Schema           `json:"items,omitempty"`
	Required   []string          `json:"required,omitempty"`
	Enum       []string          `json:"enum,omitempty"`
}

type Components struct {
	Schemas map[string]Schema `json:"schemas"`
}

type ContactInfo struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type LicenseInfo struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Helper function to extract leading comments from a location
func extractLeadingComments(method *protogen.Method) string {
	comments := method.Comments.Leading.String()
	if comments == "" {
		return ""
	}

	// Split into lines while preserving newlines
	lines := strings.Split(comments, "\n")
	var processedLines []string

	for _, line := range lines {
		// Remove comment markers but preserve indentation
		line = strings.TrimRightFunc(line, unicode.IsSpace) // Only trim right spaces
		line = strings.TrimPrefix(line, "//")
		line = strings.TrimPrefix(line, "/*")
		line = strings.TrimSuffix(line, "*/")

		// If it's an empty line (after removing markers), keep it as an empty line
		if strings.TrimSpace(line) == "" {
			processedLines = append(processedLines, "")
			continue
		}

		// If there was a comment marker, trim one leading space if it exists
		// This handles the common case of "// Comment" -> "Comment"
		if strings.HasPrefix(line, " ") {
			line = line[1:]
		}

		processedLines = append(processedLines, line)
	}

	// Join lines back together with newlines
	// Trim any leading/trailing empty lines from the final result
	return strings.Trim(strings.Join(processedLines, "\n"), "\n")
}

func main() {
	// Create a debug log file
	debugFile, err := os.Create("openapi-debug.log")
	if err != nil {
		fmt.Printf("Failed to create debug file: %v\n", err)
	} else {
		defer debugFile.Close()
	}

	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Printf("buf-plugin-openapi %s\n", "v0.1.0")
		os.Exit(0)
	}

	var flags flag.FlagSet
	options := struct {
		EmitDefaults          bool
		OmitEnumDefaultValue  bool
		EnumsAsInts           bool
		SimpleOperationIDs    bool
		UseJSONNamesForFields bool
	}{
		EmitDefaults:          true,
		OmitEnumDefaultValue:  false,
		EnumsAsInts:           false,
		SimpleOperationIDs:    false,
		UseJSONNamesForFields: true,
	}

	flags.BoolVar(&options.EmitDefaults, "emit_defaults", options.EmitDefaults, "emit default values in OpenAPI output")
	flags.BoolVar(&options.OmitEnumDefaultValue, "omit_enum_default_value", options.OmitEnumDefaultValue, "omit default value of enum types")
	flags.BoolVar(&options.EnumsAsInts, "enums_as_ints", options.EnumsAsInts, "whether to render enum values as integers")
	flags.BoolVar(&options.SimpleOperationIDs, "simple_operation_ids", options.SimpleOperationIDs, "whether to remove the service prefix from the operation IDs")
	flags.BoolVar(&options.UseJSONNamesForFields, "json_names_for_fields", options.UseJSONNamesForFields, "use json names for fields")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		// Check if we have any files to process
		hasFiles := false
		for _, f := range gen.Files {
			if f.Generate && !strings.HasPrefix(f.Desc.Path(), "google/") {
				hasFiles = true
				break
			}
		}

		// Skip if no files to process
		if !hasFiles {
			return nil
		}

		// Initialize two OpenAPI specs - one for all paths and one for public paths
		allPathsSpec := createInitialSpec("EigenLayer API Specification", "Complete API specification for the EigenLayer protocol, including all services and types.")
		publicPathsSpec := createInitialSpec("EigenLayer Public API Specification", "Public API specification for the EigenLayer protocol, containing only publicly accessible endpoints.")

		// Track unique tags for both specs
		seenTags := make(map[string]bool)
		seenPublicTags := make(map[string]bool)

		// Process all files
		for _, f := range gen.Files {
			// Skip files that are not marked for generation and google APIs
			if !f.Generate || strings.HasPrefix(f.Desc.Path(), "google/") {
				continue
			}

			// Get package name for tagging
			moduleTag := strings.ReplaceAll(string(f.Desc.Package()), ".", "/")
			if strings.HasPrefix(moduleTag, "eigenlayer/sidecar/v1/") {
				moduleTag = strings.TrimPrefix(moduleTag, "eigenlayer/sidecar/v1/")
			}

			// Add tag if not seen before
			if !seenTags[moduleTag] {
				seenTags[moduleTag] = true
				allPathsSpec.Tags = append(allPathsSpec.Tags, Tag{
					Name:        moduleTag,
					Description: fmt.Sprintf("Operations from package %s", moduleTag),
				})
			}

			// Process all services in the file
			for _, service := range f.Services {
				// Create service-specific tag
				serviceTag := fmt.Sprintf("%s/%s", moduleTag, strings.ToLower(service.GoName))
				if !seenTags[serviceTag] {
					seenTags[serviceTag] = true
					allPathsSpec.Tags = append(allPathsSpec.Tags, Tag{
						Name:        serviceTag,
						Description: fmt.Sprintf("Operations for service %s", service.GoName),
					})
				}

				for _, method := range service.Methods {
					// Get HTTP rules from annotations
					opts := method.Desc.Options()
					if opts == nil {
						continue
					}

					rule := proto.GetExtension(opts, annotations.E_Http)
					httpRule, ok := rule.(*annotations.HttpRule)
					if !ok || httpRule == nil {
						continue
					}

					// Process HTTP rule
					path, verb := extractPathAndVerb(httpRule)
					if path == "" || verb == "" {
						continue
					}

					// Create operation with module and service tags
					operation := &Operation{
						Summary:     method.GoName,
						Description: extractLeadingComments(method),
						OperationID: fmt.Sprintf("%s_%s", service.GoName, method.GoName),
						Tags:        []string{moduleTag, serviceTag},
						Responses: map[string]Response{
							"200": {
								Description: "Successful response",
								Content: map[string]MediaTypeObject{
									"application/json": {
										Schema: messageToSchema(method.Output, &options),
									},
								},
							},
						},
					}

					// Check for isPublic annotation
					isPublic := false
					if opts != nil {
						isPublicVal := proto.GetExtension(opts, openapiAnnotations.E_IsPublic)
						if isPublicVal != nil {
							isPublic = isPublicVal.(bool)
							if isPublic {
								operation.Tags = append(operation.Tags, "public_rpc")
							}
						}
					}

					// Add request body for POST/PUT/PATCH
					if verb != "get" && verb != "delete" {
						operation.RequestBody = &RequestBody{
							Required: true,
							Content: map[string]MediaTypeObject{
								"application/json": {
									Schema: messageToSchema(method.Input, &options),
								},
							},
						}
					}

					// Add path parameters
					pathParams := extractPathParams(path)
					if len(pathParams) > 0 {
						operation.Parameters = make([]Parameter, len(pathParams))
						for i, param := range pathParams {
							operation.Parameters[i] = Parameter{
								Name:     param,
								In:       "path",
								Required: true,
								Schema: Schema{
									Type: "string",
								},
							}
						}
					}

					// Add operation to paths in allPathsSpec
					pathItem := allPathsSpec.Paths[path]
					switch verb {
					case "get":
						pathItem.Get = operation
					case "post":
						pathItem.Post = operation
					case "put":
						pathItem.Put = operation
					case "delete":
						pathItem.Delete = operation
					case "patch":
						pathItem.Patch = operation
					}
					allPathsSpec.Paths[path] = pathItem

					// If the method is public, add it to publicPathsSpec
					if isPublic {
						// Add tags to publicPathsSpec if not already present
						if !seenPublicTags[moduleTag] {
							seenPublicTags[moduleTag] = true
							publicPathsSpec.Tags = append(publicPathsSpec.Tags, Tag{
								Name:        moduleTag,
								Description: fmt.Sprintf("Operations from package %s", moduleTag),
							})
						}
						if !seenPublicTags[serviceTag] {
							seenPublicTags[serviceTag] = true
							publicPathsSpec.Tags = append(publicPathsSpec.Tags, Tag{
								Name:        serviceTag,
								Description: fmt.Sprintf("Operations for service %s", service.GoName),
							})
						}

						// Add operation to publicPathsSpec
						publicPathItem := publicPathsSpec.Paths[path]
						switch verb {
						case "get":
							publicPathItem.Get = operation
						case "post":
							publicPathItem.Post = operation
						case "put":
							publicPathItem.Put = operation
						case "delete":
							publicPathItem.Delete = operation
						case "patch":
							publicPathItem.Patch = operation
						}
						publicPathsSpec.Paths[path] = publicPathItem
					}

					// Add input/output message schemas to components
					addMessageSchemas(method.Input, &allPathsSpec, &options)
					addMessageSchemas(method.Output, &allPathsSpec, &options)
					if isPublic {
						addMessageSchemas(method.Input, &publicPathsSpec, &options)
						addMessageSchemas(method.Output, &publicPathsSpec, &options)
					}
				}
			}

			// Process all messages in the file
			for _, message := range f.Messages {
				addMessageSchemas(message, &allPathsSpec, &options)
				addMessageSchemas(message, &publicPathsSpec, &options)
			}

			// Process all enums in the file
			for _, enum := range f.Enums {
				addEnumSchema(enum, &allPathsSpec, &options)
				addEnumSchema(enum, &publicPathsSpec, &options)
			}
		}

		// Write the complete spec to a file
		allData, err := json.MarshalIndent(allPathsSpec, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal complete OpenAPI spec: %v", err)
		}

		allFile := gen.NewGeneratedFile("api.swagger.json", "")
		if _, err := allFile.Write(allData); err != nil {
			return fmt.Errorf("failed to write complete OpenAPI spec: %v", err)
		}

		// Write the public spec to a file
		publicData, err := json.MarshalIndent(publicPathsSpec, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal public OpenAPI spec: %v", err)
		}

		publicFile := gen.NewGeneratedFile("api.public.swagger.json", "")
		if _, err := publicFile.Write(publicData); err != nil {
			return fmt.Errorf("failed to write public OpenAPI spec: %v", err)
		}

		return nil
	})
}

// Helper function to create an initial OpenAPI spec
func createInitialSpec(title string, description string) OpenAPISpec {
	return OpenAPISpec{
		OpenAPI: "3.0.0",
		Info: OpenAPIInfo{
			Title:       title,
			Version:     "1.0.0",
			Description: description,
			Contact: &ContactInfo{
				Name:  "EigenLayer Team",
				URL:   "https://www.eigenlayer.xyz",
				Email: "info@eigenlayer.xyz",
			},
			License: &LicenseInfo{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
		},
		Paths: make(map[string]PathItem),
		Components: &Components{
			Schemas: make(map[string]Schema),
		},
		Tags: []Tag{},
	}
}

func addMessageSchemas(message *protogen.Message, spec *OpenAPISpec, options *struct {
	EmitDefaults          bool
	OmitEnumDefaultValue  bool
	EnumsAsInts           bool
	SimpleOperationIDs    bool
	UseJSONNamesForFields bool
}) {
	// Skip if already processed
	if _, exists := spec.Components.Schemas[message.GoIdent.GoName]; exists {
		return
	}

	// Add the message schema
	spec.Components.Schemas[message.GoIdent.GoName] = messageToSchema(message, options)

	// Process nested messages
	for _, field := range message.Fields {
		if field.Message != nil && field.Message.Desc.FullName() != "google.protobuf.Timestamp" {
			addMessageSchemas(field.Message, spec, options)
		}
	}
}

func addEnumSchema(enum *protogen.Enum, spec *OpenAPISpec, options *struct {
	EmitDefaults          bool
	OmitEnumDefaultValue  bool
	EnumsAsInts           bool
	SimpleOperationIDs    bool
	UseJSONNamesForFields bool
}) {
	// Skip if already processed
	if _, exists := spec.Components.Schemas[enum.GoIdent.GoName]; exists {
		return
	}

	enumSchema := Schema{
		Type: "string",
		Enum: make([]string, len(enum.Values)),
	}

	if options.EnumsAsInts {
		enumSchema.Type = "integer"
		for i, value := range enum.Values {
			enumSchema.Enum[i] = fmt.Sprintf("%d", value.Desc.Number())
		}
	} else {
		for i, value := range enum.Values {
			enumSchema.Enum[i] = string(value.Desc.Name())
		}
	}

	spec.Components.Schemas[enum.GoIdent.GoName] = enumSchema
}

func extractPathAndVerb(rule *annotations.HttpRule) (string, string) {
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return pattern.Get, "get"
	case *annotations.HttpRule_Post:
		return pattern.Post, "post"
	case *annotations.HttpRule_Put:
		return pattern.Put, "put"
	case *annotations.HttpRule_Delete:
		return pattern.Delete, "delete"
	case *annotations.HttpRule_Patch:
		return pattern.Patch, "patch"
	}
	return "", ""
}

func extractPathParams(path string) []string {
	var params []string
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			param := part[1 : len(part)-1]
			params = append(params, param)
		}
	}
	return params
}

func messageToSchema(message *protogen.Message, options *struct {
	EmitDefaults          bool
	OmitEnumDefaultValue  bool
	EnumsAsInts           bool
	SimpleOperationIDs    bool
	UseJSONNamesForFields bool
}) Schema {
	schema := Schema{
		Type:       "object",
		Properties: make(map[string]Schema),
	}

	for _, field := range message.Fields {
		fieldSchema := fieldToSchema(field)
		if options.UseJSONNamesForFields {
			schema.Properties[field.Desc.JSONName()] = fieldSchema
		} else {
			schema.Properties[string(field.Desc.Name())] = fieldSchema
		}

		if field.Desc.HasPresence() {
			schema.Required = append(schema.Required, string(field.Desc.Name()))
		}
	}

	return schema
}

func fieldToSchema(field *protogen.Field) Schema {
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return Schema{Type: "boolean"}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return Schema{Type: "integer", Format: "int32"}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return Schema{Type: "integer", Format: "int64"}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return Schema{Type: "integer", Format: "uint32"}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return Schema{Type: "integer", Format: "uint64"}
	case protoreflect.FloatKind:
		return Schema{Type: "number", Format: "float"}
	case protoreflect.DoubleKind:
		return Schema{Type: "number", Format: "double"}
	case protoreflect.StringKind:
		return Schema{Type: "string"}
	case protoreflect.BytesKind:
		return Schema{Type: "string", Format: "byte"}
	case protoreflect.EnumKind:
		return Schema{Type: "string"}
	case protoreflect.MessageKind:
		if field.Message.Desc.FullName() == "google.protobuf.Timestamp" {
			return Schema{Type: "string", Format: "date-time"}
		}
		return Schema{
			Ref: fmt.Sprintf("#/components/schemas/%s", field.Message.GoIdent.GoName),
		}
	}
	return Schema{Type: "string"}
}
