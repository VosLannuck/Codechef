package beginnerlevel
// https://www.codechef.com/problems/THREETOPICS
import "fmt"
func main () {

  var a, b, c , x int;
  fmt.Scan(&a, &b, &c, &x);
  if(a == x || b == x || c == x ) {
    fmt.Println("YES");
  } else {
    fmt.Println("NO");
  }

}
