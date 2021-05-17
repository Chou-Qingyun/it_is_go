// echo1 输出其命令行参数
/*package main
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}*/

// echo2 输出其命令行参数
/*package main
import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}*/

// echo3 输出其命令行参数
package main
import (
	"fmt"
	"os"
	"strings"
)
func main() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "));
}