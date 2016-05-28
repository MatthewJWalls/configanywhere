package formats

import (
	"strings"
	"strconv"
	"reflect"
)

type KeyValueFormat struct {
	target interface{}
}

func (this KeyValueFormat) Using(bytes []byte) {

	str := string(bytes)

	for _, e := range(strings.Split(str, "\n")) {

		if ! strings.Contains(e, "=") {
			continue
		}

		parts := strings.Split(e, "=")
		key := parts[0]
		val := parts[1]

		te := reflect.ValueOf(this.target).Elem()

		for i := 0; i < te.NumField(); i++ {

			field := te.Field(i)
			typ := te.Type().Field(i)

			if typ.Tag.Get("key") == key && field.CanSet() {

				if field.Kind() == reflect.Int {
					// set an int
					intval, _ := strconv.Atoi(val)
					field.SetInt(int64(intval))
				} else if field.Kind() == reflect.Float64 {
					// set a float
					floatval, _ := strconv.ParseFloat(val, 64)
					field.SetFloat(float64(floatval))
				} else if field.Kind() == reflect.Bool {
					// set a boolean
					field.SetBool(val == "true")
				} else {
					// set a String
					field.SetString(string(val))
				}
				
			}
			
		}
				
	}
	
}

func NewKeyValueFormat(target interface{}) KeyValueFormat {
	return KeyValueFormat{ target }
}
