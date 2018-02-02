package mysqlx

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type decimalPair struct {
	b        []byte
	expected string
	valid    bool
}

// values from TestQueryData
var decimalPairs = []decimalPair{
	{[]byte{0x00, 0x9c}, "9", true},
	{[]byte{0x00, 0x9d}, "-9", true},
	{[]byte{0x00, 0x12, 0xc0}, "12", true},
	{[]byte{0x00, 0x12, 0xd0}, "-12", true},
	{[]byte{0x01, 0x09, 0xc0}, "0.9", true},
	{[]byte{0x01, 0x09, 0xd0}, "-0.9", true},
	{[]byte{0x03, 0x12, 0x34, 0x0c}, "12.340", true},
	{[]byte{0x03, 0x12, 0x34, 0x0d}, "-12.340", true},
	{[]byte{0x04, 0x12, 0x34, 0x01, 0xc0}, "12.3401", true},
	{[]byte{0x04, 0x12, 0x34, 0x01, 0xd0}, "-12.3401", true},

	{nil, "", false},
	{[]byte{}, "", false},
	{[]byte{0x00}, "", false},
	{[]byte{0x00, 0x00}, "", false},
	{[]byte{0x30, 0x30}, "", false},
	{[]byte{0xff, 0xff}, "", false},
	{[]byte{0x30, 0x0a, 0x30}, "", false},
}

func TestUnmarshalDecimal(t *testing.T) {
	t.Parallel()

	for _, p := range decimalPairs {
		d, err := unmarshalDecimal(p.b)
		assert.Equal(t, p.valid, err == nil)
		assert.Equal(t, p.expected, d)
	}
}

func TestUnmarshalDecimalFuzzCorpus(t *testing.T) {
	path := filepath.Join("fuzz", "UnmarshalDecimal", "corpus")
	err := os.MkdirAll(path, 0777)
	require.NoError(t, err)
	for i, p := range decimalPairs {
		err := ioutil.WriteFile(filepath.Join(path, fmt.Sprintf("test-%d", i)), p.b, 0666)
		require.NoError(t, err)
	}
}

var sink interface{}

func BenchmarkUnmarshalDecimal(b *testing.B) {
	for _, p := range decimalPairs {
		b.Run(p.expected, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink, sink = unmarshalDecimal(p.b)
			}
		})
	}
}
