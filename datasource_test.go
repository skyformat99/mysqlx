package mysqlx

import (
	"crypto/tls"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDataSource(t *testing.T) {
	t.Parallel()

	for s, expected := range map[string]*DataSource{
		"mysqlx://my_user:my_password@127.0.0.1:33060/world_x?_tls-insecure=true&time_zone=UTC": &DataSource{
			Host:     "127.0.0.1",
			Port:     33060,
			Database: "world_x",
			Username: "my_user",
			Password: "my_password",
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			SessionVariables: map[string]string{"time_zone": "UTC"},
		},
	} {
		t.Run(s, func(t *testing.T) {
			t.Parallel()

			u, err := url.Parse(s)
			require.NoError(t, err)
			actual, err := ParseDataSource(u)
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
			assert.Equal(t, u, actual.URL(), "%s")
		})
	}
}
