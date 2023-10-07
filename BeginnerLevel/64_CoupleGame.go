// https://www.codechef.com/problems/COUGAME
package beginnerlevel
import "fmt"

func main() {
  var t,g,b int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&g, &b);
    fmt.Println(b-g);
    t--;
  }
}
