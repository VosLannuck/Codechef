package beginnerlevel
// https://www.codechef.com/problems/PIZZAC
import "fmt"

func main() {

  var t, n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n);
    if( n % 2== 0|| n == 1) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }

}
