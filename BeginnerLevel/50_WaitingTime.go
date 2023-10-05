package beginnerlevel
// https://www.codechef.com/problems/WAITTIME
import "fmt"

func main() {
  
  var t,k,x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&k, &x);
    fmt.Println(k * 7 - x);
    t--;
  }
}
