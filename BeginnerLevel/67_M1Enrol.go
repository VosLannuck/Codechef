// https://www.codechef.com/problems/M1ENROL
package beginnerlevel
import "fmt"

func main() {
  var t, x, y int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y);
    if( y <= x) {
      fmt.Println(0);
    } else {
      fmt.Println(y - x);
    }
    t--;
  }

}
