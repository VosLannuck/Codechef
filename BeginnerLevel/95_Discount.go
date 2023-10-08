// https://www.codechef.com/problems/DISCNT
package beginnerlevel
import "fmt"
func main() {
  var t, x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    var discount int = 100 * x / 100;
    fmt.Println(100 - discount);
    t--;
  } 
}
