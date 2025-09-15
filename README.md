# enchantedeland

CUE playground for Go→YAML marshaling

## Usage

```bash
go run main.go
```

## What it does

- Takes Go data structures
- Converts via CUE
- Outputs YAML

## Code

```go
yamlString, err := YAMLMarshalWithCUE(data)
```
