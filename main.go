package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

// YAMLMarshalWithCUE converts a Go data structure to YAML string using CUE
func YAMLMarshalWithCUE(data interface{}) (string, error) {
	// Create a new CUE context
	ctx := cuecontext.New()
	
	// Encode the Go data structure into a CUE Value
	value := ctx.Encode(data)
	
	// Check for encoding errors
	if err := value.Err(); err != nil {
		return "", fmt.Errorf("failed to encode data to CUE: %w", err)
	}
	
	// Convert the CUE Value to YAML bytes
	yamlBytes, err := yaml.Encode(value)
	if err != nil {
		return "", fmt.Errorf("failed to encode CUE value to YAML: %w", err)
	}
	
	return string(yamlBytes), nil
}

// Example usage
func main() {
	// Example Go data structure
	data := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"active":  true,
		"hobbies": []string{"reading", "coding", "hiking"},
		"address": map[string]string{
			"street": "123 Main St",
			"city":   "Anytown",
			"zip":    "12345",
		},
	}
	
	yamlString, err := YAMLMarshalWithCUE(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	
	fmt.Println("YAML output:")
	fmt.Println(yamlString)
}
