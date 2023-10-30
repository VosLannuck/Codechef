package beginnerlevel
// https://www.codechef.com/problems/CHEFCAND
import "fmt"
import "math"
func main() {

    var t int;
  var  n, x float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x);
    if(x > n ) {
      fmt.Println(0);
    } else{
   
    fmt.Println( math.Ceil((n - x ) / 4));
    }
    t--;
  }
}
