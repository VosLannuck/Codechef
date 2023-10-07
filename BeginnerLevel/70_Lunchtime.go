package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/LTIME

func main() {

  var t,x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    if(x < 1 || x > 4) {
      fmt.Println("NO");
    } else {
      fmt.Println("YeS");
    }
    t--;
  }
}
