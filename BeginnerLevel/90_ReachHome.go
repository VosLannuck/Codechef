package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/REACH_HOME
func main() {

  var t,x, y int;
  fmt.Scan(&t);
  for t >0 {
    fmt.Scan(&x, &y);
    if(x * 5 >= y) {
      fmt.Println("Yes");
    } else {
      fmt.Println("No");
    }
    t--;
  }

}
