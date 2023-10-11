// https://www.codechef.com/problems/FLOW002
package beginnerlevel
import "fmt"
var t, a, b int;

func main() {

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b);
    fmt.Println(a % b);
    t--;
  }

}
