package generator

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"reciper/recipe-generator/internal/domain"
)

// 4d17b5d37a11046bfcddeb6206ab2f0380651fdd36ebd9995e9b8f3dfe65dce6
// 4d1 - prefix
// 7b5 - adjective
// d3 - icon
// 7a 110 | 46 bfc| dd eb6 | 20 6ab | 2f 038 | 06 51f | dd 36e | bd 999 | 5e 9b8 | f3 dfe | 65 dce6

func GenerateRecipe(username string) domain.Recipe {
	hashBytes := sha256.Sum256([]byte(username))
	hash := hex.EncodeToString(hashBytes[:])
	prefix := getPrefix(hexToInt(hash[:3]))
	adjective := getAdjective(hexToInt(hash[3:6]))
	icon := hexToInt(hash[6:8])
	ingredients := []domain.Ingredient{}
	ingredients = append(ingredients, getFirstIngredient(hexToInt(hash[8:10]), hexToInt(hash[10:13])))
	for i := 1; i < 11; i++ {
		step := i * 5
		ingredient := getIngredient(resolveQuantity(hexToInt(hash[8+step:10+step])), hexToInt(hash[10+step:13+step]))
		if ingredient.Quantity == 0 {
			continue
		}
		ingredients = append(ingredients, ingredient)
	}
	return domain.Recipe{
		Hash:        hash,
		Prefix:      prefix,
		Adjective:   adjective,
		Icon:        icon,
		Ingredients: ingredients,
	}
}

func hexToInt(hex string) int {
	num, _ := strconv.ParseInt(hex, 16, 64)
	return int(num)
}

func getPrefix(prefixPiece int) string {
	prefixes := GetAsset().Prefixes
	return prefixes[prefixPiece%len(prefixes)]
}

func getAdjective(adjectivePiece int) string {
	adjectives := GetAsset().Adjectives
	return adjectives[adjectivePiece%len(adjectives)]
}

func getFirstIngredient(quantityPiece int, ingredientPiece int) domain.Ingredient {
	var quantity int
	if quantityPiece <= 200 {
		quantity = 1
	} else {
		quantity = 2
	}
	return getIngredient(quantity, ingredientPiece)
}

func getIngredient(quantity int, ingredientPiece int) domain.Ingredient {
	ingredients := GetAsset().Ingredients
	ingredient := ingredients[ingredientPiece%len(ingredients)]
	return domain.Ingredient{
		Name:     ingredient,
		Quantity: uint8(quantity),
	}
}

func resolveQuantity(piece int) int {
	var quantity int
	switch {
	case piece <= 100:
		quantity = 0
	case piece > 100 && piece < 170:
		quantity = 1
	case piece >= 170 && piece < 255:
		quantity = 2
	case piece == 255:
		quantity = 3
	}
	return quantity
}
