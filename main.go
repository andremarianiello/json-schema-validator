package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

func main() {

	schemaPath, err := filepath.Abs(os.Args[1])
	if err != nil {
		panic(err)
	}
	documentPath, err := filepath.Abs(os.Args[2])
	if err != nil {
		panic(err)
	}
	schemaLoader := gojsonschema.NewSchemaLoader()
	schemaJsonLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentJsonLoader := gojsonschema.NewReferenceLoader("file://" + documentPath)
	schema, err := schemaLoader.Compile(schemaJsonLoader)
	if err != nil {
		panic(err.Error())
	}

	result, err := schema.Validate(documentJsonLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		os.Exit(1)
	}
}
