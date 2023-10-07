// https://www.codechef.com/problems/PARTY2
package beginnerlevel
import "fmt"

func main() {
  var t, n, x, k int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x, &k);
    if(n * x <= k) {
      fmt.Println("Yes");
    } else {
      fmt.Println("No");
    }
    t--;
  }
}
