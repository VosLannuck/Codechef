package beginnerlevel
// https://www.codechef.com/problems/INSTAGRAM
import "fmt"
func main() {

  var t int;
  fmt.Scan(&t);
  for t > 0 {
    var x, y int;
    fmt.Scan(&x, &y);
    if( x > 10 * y) {
      fmt.Println("yes");
    }  else {
      fmt.Println("no");
    }
    t--;
  }

}
