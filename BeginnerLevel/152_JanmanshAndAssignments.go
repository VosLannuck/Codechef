package beginnerlevel
// https://www.codechef.com/problems/JASSIGNMENTS
import "fmt"

func main() {

  var t, x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
  if( (x + 3) <= 10) {
    fmt.Println("YES");
  } else {
    fmt.Println("NO");
  }
  
    t--;
  }
}
