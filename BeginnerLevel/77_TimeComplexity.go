package beginnerlevel
//https://www.codechef.com/problems/COMPLEXITY 
import "fmt"

func main() {
  var t, x, y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if( x > y) {
      fmt.Println("Yes");
    } else {
      fmt.Println("No");
    }
    t--;
  }

}
