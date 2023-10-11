package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/EXPIRY
func main() {

  var t int ;
  var n,m,k float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n,&m, &k);
    var total = n / k ;
    if( total <= m) {
      fmt.Println("YES");
    } else {
      fmt.Println("no");
    }
    t--;
  }

}

