package generator

import (
	"encoding/json"
	"reciper/recipe-generator/internal/generator/assets"
)

var asset *Asset

type Asset struct {
	Prefixes    []string `json:"prefixes"`    // 66
	Adjectives  []string `json:"adjectives"`  // 457
	Ingredients []string `json:"ingredients"` // 380
}

func GetAsset() Asset {
	if asset == nil {
		getAsset()
	}
	return *asset

}

func getAsset() {
	json.Unmarshal(assets.Json, &asset)

}
