// https://www.codechef.com/problems/TYRE

package beginnerlevel
import "fmt"

func main() {
  var t, n, m int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &m);
    fmt.Println(n * 2 + 4 * m);
    t--;
  }
}
