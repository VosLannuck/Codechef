package beginnerlevel
// https://www.codechef.com/problems/TEA
import "fmt"

func main() {

  var t, x ,y,z int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y, &z);
    if( x % y == 0) {
      fmt.Println((x / y) *z );
    } else {
      fmt.Println( (( x / y ) + 1 )  *z);
    }
    t--;
  }
  

}
