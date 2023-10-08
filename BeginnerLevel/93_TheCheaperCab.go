package beginnerlevel
// https://www.codechef.com/problems/CABS
import "fmt"
func main() {

  var t, x, y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if(x == y) {
      fmt.Println("Any");
    } else if (x > y) {
      fmt.Println("second");
    } else {
      fmt.Println("first");
    }
    t--;
  }

}
