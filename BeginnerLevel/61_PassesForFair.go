package beginnerlevel
// https://www.codechef.com/problems/FAIRPASS
import "fmt"
func main() {

  var t, n, k int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &k)
    if(k > n) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }

}
