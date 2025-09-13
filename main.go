package main

import (
	"encoding/json"
	"fmt"
)
import "sera/ops/domain"

func main() {
	meta := domain.CreateMeta()
	b, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
