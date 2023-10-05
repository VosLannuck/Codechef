// https://www.codechef.com/problems/SUBSCRIBE

package beginnerlevel
import "fmt"

func main() {
  var test, x int;

  fmt.Scan(&test);

  for test > 0 {
    fmt.Scan(&x);
    if(x > 30) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }

    test--;
  }

}
