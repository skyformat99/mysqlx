package mysqlx

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/url"
)

// Driver implements database/sql/driver.Driver interface.
type Driver struct{}

// Open returns a new connection to the database. See README for dataSource format.
// The returned connection may be used only by one goroutine at a time.
func (Driver) Open(dataSource string) (driver.Conn, error) {
	u, err := url.Parse(dataSource)
	if err != nil {
		return nil, err
	}
	ds, err := ParseDataSource(u)
	if err != nil {
		return nil, err
	}
	return open(context.Background(), ds)
}

// OpenCtx is a variant of Open with specified context and DataSource struct.
// Context is used only for connection establishing: dialing, negotiating
// and authenticating. Canceling the context after the connection is established does nothing.
func (Driver) OpenCtx(ctx context.Context, dataSource *DataSource) (driver.Conn, error) {
	return open(ctx, dataSource)
}

func init() {
	sql.Register("mysqlx", Driver{})
}

// check interfaces
var (
	_ driver.Driver = Driver{}
)
