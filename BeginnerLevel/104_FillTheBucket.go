// https://www.codechef.com/problems/FBC
package beginnerlevel
import "fmt"

func main() {
  var t,x ,k int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&k, &x);
    fmt.Println(k - x);
    t--;
  }
}
