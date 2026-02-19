package main

import "fmt"

func formatProxy(ip, port, user, pass string) string {
	return fmt.Sprintf("http://%s:%s@%s:%s", user, pass, ip, port)
}
