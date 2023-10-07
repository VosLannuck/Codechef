package beginnerlevel
// https://www.codechef.com/problems/APPLORNG
import "fmt"

func main() {
  
  var x, a, b int;
  fmt.Scan(&x);
  fmt.Scan(&a, &b);
  if(a + b <= x) {
    fmt.Println("YES");
  } else {
    fmt.Println("No");
  }
}
