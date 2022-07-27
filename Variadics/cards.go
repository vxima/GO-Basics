package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	s := []int{2, 6, 9}
	return s
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	} else {
		return slice[index]
	}

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	size := len(slice)
	if index >= size || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	aux := []int{}
	aux = append(aux, values...)
	aux = append(aux, slice...)
	return aux
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
		return slice
	} else {
		preIndex := []int{}
		for i, v := range slice {
			if i == index {
				break
			} else {
				preIndex = append(preIndex, v)
			}
		}
		for i, v := range slice {
			if i > index {
				preIndex = append(preIndex, v)
			}
		}
		return preIndex
	}
}
