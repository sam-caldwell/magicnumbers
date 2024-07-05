package magicnumbers

// GetLongest - return the longest format in the Map
func (m *Map) GetLongest() (result int) {
	for _, v := range *m {
		if sz := len(v); sz > result {
			result = sz
		}
	}
	return result
}
