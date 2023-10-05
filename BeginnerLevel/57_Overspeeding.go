package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/FINE
func main() {

  var t, x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    if(x > 70 && x <= 100) {
      fmt.Println(500);
    } else if (x > 100) {
      fmt.Println(2000);
    } else {
      fmt.Println(0);
    }
    t--;
  }
}
