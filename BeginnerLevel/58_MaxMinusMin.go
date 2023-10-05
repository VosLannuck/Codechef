package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/MAXDIFFMIN
func main() {

  var t, a, b, c int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&a, &b, &c);
    fmt.Println(c - a);
    t--;
  }
  
}

