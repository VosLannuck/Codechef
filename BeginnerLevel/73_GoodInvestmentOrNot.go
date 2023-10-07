package beginnerlevel
// https://www.codechef.com/problems/INVESTMENT
import "fmt"

func main() {
  var t,x,y int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if(y * 2 > x ) {
      fmt.Println("No");
    } else {
      fmt.Println("Yes");
    }
    t--;
  }

}
