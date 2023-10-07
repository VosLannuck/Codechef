package beginnerlevel
// https://www.codechef.com/problems/WATERREQ
import "fmt"

func main() {
  var t,n int;
  fmt.Scan(&t);
  for t > 0 {
    
    fmt.Scan(&n);  
    fmt.Println(2 * n);
    t--;
  }
}
