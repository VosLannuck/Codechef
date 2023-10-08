package beginnerlevel
// https://www.codechef.com/problems/CWIREFRAME
import "fmt"

func main() {

  var t,n,m,x int;
  fmt.Scan(&t); 
  for t > 0 {
    fmt.Scan(&n, &m, &x);
    fmt.Println( (2 * (n + m)) * x);
    t--;
  } 
}
