package beginnerlevel
// https://www.codechef.com/problems/SHOPCHANGE
import "fmt"

func main() {
  
  var t , x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    fmt.Println(100-x );
    t--;
  }

}
