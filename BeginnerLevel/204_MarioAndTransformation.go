package beginnerlevel
// https://www.codechef.com/problems/TRANSFORM
import "fmt"
func main() {

  var t , x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    if(x % 3 == 0) {
        fmt.Println("NORMAL");
    } else if(x % 3 == 1) {
        fmt.Println("HUGE");
    } else {
      fmt.Println("SMALL");
    }
    t--;
  }

}
