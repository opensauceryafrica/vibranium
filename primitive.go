package vibranium

type Array []interface{}

// Includes determines whether an array includes a certain value among its entries, returning true or false as appropriate.
func (a Array) Includes(value interface{}) bool {
	for _, v := range a {
		if v == value {
			return true
		}
	}
	return false
}
