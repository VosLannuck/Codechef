// https://www.codechef.com/problems/BROKENPHONE

package beginnerlevel
import  "fmt"

func main() {
  var t,x, y int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if(x == y) {
      fmt.Println("ANY");
    } else if(x <  y)  {
      fmt.Println("Repair");
    } else {
      fmt.Println("New Phone");
    }
    t--;
  }
}
