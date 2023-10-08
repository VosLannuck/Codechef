package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/HOTCOLD
func main() {
  var t , c int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&c);
    if (c > 20 ) {
      fmt.Println("HOT");
    } else {
      fmt.Println("COLD");
    }
    t--;
  }
}
