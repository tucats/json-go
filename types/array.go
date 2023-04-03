package types

// Create a definition for an array element. If all the members of the array are
// the same type, then that array type is returned. If the array is heterogeneous,
// then it is declared as []interface{}.
func array(actual []interface{}, depth int, baseType string, camel bool, typeMap map[string]string) string {
	var t string

	for index, value := range actual {
		if index == 0 {
			t = Element(value, depth+1, baseType, camel, typeMap)
		} else {
			if t != Element(value, depth+1, baseType, camel, typeMap) {
				return "[]interface{}"
			}
		}
	}

	return "[]" + t
}
