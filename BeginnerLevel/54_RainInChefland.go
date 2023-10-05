package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/RAINFALL1
func main() {
  var t, x int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x);
      if(x < 3) {
      fmt.Println("LIGHT");
      } else if (x >= 3 && < 7) {
      fmt.Println("MODERATE");
    } else {
      fmt.Println("HEAVY");
    }
    t--;
  }
}
