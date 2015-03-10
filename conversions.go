package main

// mergeKeys will merge the mapval fields, overwriting y.mapval
// it is safe to call even if y.mapval is nil
func (y *yySymType) mergeKeys(other map[string]interface{}) {
	map1 := y.mapval
	y.mapval = mergeKeys(map1, other)
}

// mergeKeys will produce the union of two sets
// The behavior for duplicate keys is undefined
func mergeKeys(a, b map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, m := range []map[string]interface{}{a, b} {
		for k, val := range m {
			result[k] = val
		}
	}
	return result
}

func parseIdentifier(y yySymType) interface{} {
	switch y.val {
	case "true":
		return true
	case "false":
		return false
	case "null":
		return nil
	default:
		return y.val
	}
	return y
}
