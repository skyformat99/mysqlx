// +build gofuzz

package mysqlx

func FuzzUnmarshalDecimal(data []byte) int {
	d, err := unmarshalDecimal(data)
	if err != nil {
		return 0
	}
	return 1
}
