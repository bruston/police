package police

import "reflect"

// Makes a map[string]string from a structs string fields
func structToMap(i interface{}) map[string]string {
	iValue := reflect.ValueOf(i)
	t := iValue.Type()
	m := make(map[string]string)
	for i := 0; i < iValue.NumField(); i++ {
		field := iValue.Field(i)
		if field.Kind() != reflect.String {
			continue
		}
		v := field.String()
		if v != "" {
			m[t.Field(i).Name] = v
		}
	}
	return m
}
