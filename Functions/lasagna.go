package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time

}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendList, myList []string) []string {
	idx_f := len(friendList) - 1
	idx_m := len(myList) - 1
	myList[idx_m] = friendList[idx_f]
	return myList

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, quant int) []float64 {
	factor := float64(quant)
	factor = factor / 2.0
	size := len(recipe)
	newRecipe := make([]float64, size)
	for i, v := range recipe {
		newRecipe[i] = v * factor
	}
	return newRecipe
}
