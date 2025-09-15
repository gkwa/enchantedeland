package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

// ValidateWithCUE validates a Go data structure using CUE
func ValidateWithCUE(data any) error {
	ctx := cuecontext.New()
	value := ctx.Encode(data)

	if err := value.Err(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Additional validation can be added here
	// e.g., value.Validate() for more strict checking
	return nil
}

// ConvertToYAML converts a CUE value to YAML string
func ConvertToYAML(data any) (string, error) {
	ctx := cuecontext.New()
	value := ctx.Encode(data)

	if err := value.Err(); err != nil {
		return "", fmt.Errorf("failed to encode data to CUE: %w", err)
	}

	yamlBytes, err := yaml.Encode(value)
	if err != nil {
		return "", fmt.Errorf("failed to encode CUE value to YAML: %w", err)
	}

	return string(yamlBytes), nil
}

// YAMLMarshalWithCUE combines validation and conversion
func YAMLMarshalWithCUE(data any) (string, error) {
	if err := ValidateWithCUE(data); err != nil {
		return "", err
	}

	return ConvertToYAML(data)
}

// Example usage
func main() {
	data := map[string]any{
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

	// Option 1: Separate validation and conversion
	if err := ValidateWithCUE(data); err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	yamlString, err := ConvertToYAML(data)
	if err != nil {
		log.Fatalf("Conversion error: %v", err)
	}

	// Option 2: Combined validation and conversion
	// yamlString, err := YAMLMarshalWithCUE(data)
	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }

	fmt.Println("YAML output:")
	fmt.Println(yamlString)
}
