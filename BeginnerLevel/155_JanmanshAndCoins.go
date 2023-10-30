package beginnerlevel
// https://www.codechef.com/problems/JCOINS
import "fmt"

func main() {

  var t, x , y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    fmt.Println(x * 10 + y * 5);
    t--;
  }

}
