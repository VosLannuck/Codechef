package beginnerlevel
//https://www.codechef.com/problems/AIRHOCKEY
import "fmt"

func main() {
  var t, a, b int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b);
    if(a > b) {
      fmt.Println(7 - a);
    }else {
      fmt.Println( 7 - b)
    }
    t--;
  }
}
