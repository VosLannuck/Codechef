package beginnerlevel
import "fmt"
import "math"
// https://www.codechef.com/problems/VOLCONTROL
func main() {
  var t int;
  var x, y float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    fmt.Println(math.Abs( x - y));
    t--;  
  }

}

