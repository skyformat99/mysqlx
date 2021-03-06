package mysqlx

import (
	"context"
	"crypto/tls"
	"database/sql/driver"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type DataSource struct {
	Host             string
	Port             uint16
	Database         string
	Username         string
	Password         string
	SessionVariables map[string]string

	TLSConfig *tls.Config

	Trace TraceFunc
}

func (ds *DataSource) Connect(ctx context.Context) (driver.Conn, error) {
	return open(ctx, ds)
}

func (ds *DataSource) Driver() driver.Driver {
	return Driver
}

func ParseDataSource(u *url.URL) (*DataSource, error) {
	if u.Scheme != "mysqlx" {
		return nil, fmt.Errorf("unexpected scheme %s", u.Scheme)
	}
	ds := &DataSource{
		Host:     u.Hostname(),
		Database: strings.TrimPrefix(u.Path, "/"),
		Trace:    noTrace,
	}

	// set port if given
	if p := u.Port(); p != "" {
		pp, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			return nil, err
		}
		ds.Port = uint16(pp)
	}

	// set username and password if they are given
	if u.User != nil {
		ds.Username = u.User.Username()
		ds.Password, _ = u.User.Password()
	}

	for k, vs := range u.Query() {
		if len(vs) != 1 {
			return nil, fmt.Errorf("%d values given for session variable %s: %v", len(vs), k, vs)
		}
		v := vs[0]

		// set session variables
		if !strings.HasPrefix(k, "_") {
			if ds.SessionVariables == nil {
				ds.SessionVariables = make(map[string]string)
			}
			ds.SessionVariables[k] = v
			continue
		}

		switch k {
		case "_tls":
			enable, err := strconv.ParseBool(v)
			if err != nil {
				return nil, err
			}
			if enable {
				ds.TLSConfig = &tls.Config{
					ServerName: ds.Host,
				}
			}

		case "_tls-insecure":
			insecure, err := strconv.ParseBool(v)
			if err != nil {
				return nil, err
			}
			if insecure {
				ds.TLSConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}

		case "_trace":
			ds.Trace = getTracef(v)

		default:
			return nil, fmt.Errorf("unexpected parameter %q", k)
		}
	}

	return ds, nil
}

func (ds *DataSource) hostPort() string {
	return net.JoinHostPort(ds.Host, strconv.FormatUint(uint64(ds.Port), 10))
}

func (ds *DataSource) URL() *url.URL {
	u := &url.URL{
		Scheme: "mysqlx",
		Host:   ds.hostPort(),
		Path:   "/" + ds.Database,
	}

	if ds.Username != "" {
		u.User = url.UserPassword(ds.Username, ds.Password)
	}

	q := make(url.Values)
	for k, v := range ds.SessionVariables {
		q.Set(k, v)
	}

	if ds.TLSConfig != nil {
		if ds.TLSConfig.InsecureSkipVerify {
			q.Set("_tls-insecure", "true")
		} else {
			q.Set("_tls", "true")
		}
	}

	u.RawQuery = q.Encode()
	return u
}

// check interfaces
var (
// TODO _ driver.Connector = (*Connector)(nil)
)
