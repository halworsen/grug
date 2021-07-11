package grug

import "fmt"

// anything to string
func atostr(a interface{}) string {
	return fmt.Sprintf("%v", a)
}
