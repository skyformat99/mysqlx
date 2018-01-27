// +build mysqlxnopprof

package mysqlx

func connectionCreated(local, remote string) {}
func connectionClosed(local, remote string)  {}
func assertOpenConnections(expected int, failer func())
