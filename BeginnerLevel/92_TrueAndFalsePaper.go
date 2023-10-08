package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/TFPAPER
func main() {
  var t,n,k int;
  fmt.Scan(&t);
  for  t > 0 {
    fmt.Scan(&n, &k);
    fmt.Println( (n - k) );
    t--;
  }
}
