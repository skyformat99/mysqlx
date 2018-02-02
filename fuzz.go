// +build gofuzz

package mysqlx

func FuzzUnmarshalDecimal(data []byte) int {
	_, err := unmarshalDecimal(data)
	if err != nil {
		return 0
	}
	return 1
}
