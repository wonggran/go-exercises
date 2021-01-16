package flatten

func Flatten(arr interface{}) []interface{} {
	var ret []interface{}

	for _, v := range arr.([]interface{}) {
		switch v.(type) {
		case nil:
		case []interface{}:
			for _, e := range Flatten(v) {
				ret = append(ret, e)
			}
		default:
			ret = append(ret, v)
		}
	}

	return ret
}
