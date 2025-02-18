package camell

import (
	//"fmt"
	"strings"
)

func main() {

}


func IsCamellEmployee(employee string) bool {
	return strings.HasSuffix(employee, "@camell.com")
}