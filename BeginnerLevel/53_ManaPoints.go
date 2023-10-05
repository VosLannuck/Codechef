package beginnerlevel
// https://www.codechef.com/problems/MANAPTS
import "fmt"

func main() {
  var t, x,y int;
  
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    fmt.Println(y / x);
    t--;
  }

}
