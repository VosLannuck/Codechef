package beginnerlevel
// https://www.codechef.com/problems/ONEFULPAIRS
import "fmt"

func main() {

  var a, b int;
  fmt.Scan(&a, &b);

  if( (a*b ) + a + b == 111) {
    fmt.Println("YES");
  } else {
    fmt.Println("NO");
  }
  
}
