package beginnerlevel
// https://www.codechef.com/problems/CARTRIP
import "fmt"

func main() {

  var t,x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    if(x <= 300) {
      fmt.Println(300 * 10);
    } else {
      fmt.Println((x * 10));
    }
    t--;
  }


}
