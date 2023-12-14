package beginnerlevel
// https://www.codechef.com/problems/CANDYDIST
import "fmt"

func main() {

  var t, n, m int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &m);
    var distributedCandies = n / m;
    if(distributedCandies % 2 == 0 && n % m == 0) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }
}
