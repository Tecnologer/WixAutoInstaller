package main

import (
	"flag"
	"fmt"
)

var product = flag.String("p", "test/test_product.txt", "Wix product file")
var assembly = flag.String("a", "test/test_assembly.txt", "Assembly info file")
var version = flag.String("v", "0.0.1.004", "Version for the new installer")

func main() {
	flag.Parse()

	err := UpdateProduct(*product, *version)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %s;\nUpdated with version: %s\n", *product, *version)

	err = UpdateAssembly(*assembly, *version)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Assembly Info: %s;\nUpdated with version: %s\n", *assembly, *version)
}
