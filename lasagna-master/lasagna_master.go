package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (gramsOfNoodles int, litersOfSauce float64) {
	const noodlesPerLayer = 50
	const saucePerLayer = 0.2

	for _, ingredient := range layers {
		switch ingredient {
		case "noodles":
			gramsOfNoodles += noodlesPerLayer
		case "sauce":
			litersOfSauce += saucePerLayer
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	const portionsInRecipe = 2
	updatedQuantities := make([]float64, len(quantities))

	for i, amount := range quantities {
		updatedQuantities[i] = amount * float64(portions) / portionsInRecipe
	}

	return updatedQuantities
}
