// +build !mysqlxnopprof

package mysqlx

import (
	"fmt"
	"os"
	"runtime/pprof"
)

var (
	connectionsProfile = pprof.NewProfile("github.com/AlekSi/mysqlx.connections")
)

func connectionCreated(local, remote string) {
	key := fmt.Sprintf("%s-%s", local, remote)
	connectionsProfile.Add(key, 1)
}

func connectionClosed(local, remote string) {
	key := fmt.Sprintf("%s-%s", local, remote)
	connectionsProfile.Remove(key)
}

func assertOpenConnections(expected int, failer func()) {
	actual := connectionsProfile.Count()
	if expected != actual {
		fmt.Fprintf(os.Stderr, "expected %d connections, have %d:\n", expected, actual)
		if err := connectionsProfile.WriteTo(os.Stderr, 1); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		failer()
	}
}
