// https://www.codechef.com/problems/PARLIAMENT
package beginnerlevel
import "fmt"

func main () {
  var t int;
  var n , x float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n ,&x);
    if(x >= n/2 ) {
      fmt.Println("YES");
    } else {
      fmt.Println("No");
    }
    t--;
  }

}
