package types

import "math"

// element converts a single element of JSON. This can be a scalar item like a
// string or boolean, or it might be an array or structure. It can call itself
// recursively to process the array elements and structure fields.
func Element(x interface{}, depth int, baseType string, camel bool, typeMap map[string]string) string {
	switch actual := x.(type) {
	case map[string]interface{}:
		return structure(actual, depth, baseType, camel, typeMap)

	case bool:
		return "bool"

	case int:
		return "int"

	case int64:
		return "int64"

	case string:
		return "string"

	case float64:
		if actual == math.Floor(actual) {
			return "int"
		}

		return "float64"

	case []interface{}:
		return array(actual, depth, baseType, camel, typeMap)

	default:
		return "<unknown>"
	}
}
