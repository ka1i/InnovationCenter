package usage

import (
	"fmt"
)

// Usage ...
func Usage() {
	fmt.Println(`Usage: InnovationCenter -[hv] 

If you are using it for the first time,
first execute the command "InnovationCenter -h" to check usage.

	 ------- < Commands Arguments > -------
optional:
  -h, help		Show this help message. 
  -v, version		Show the app version.
  `)
}
