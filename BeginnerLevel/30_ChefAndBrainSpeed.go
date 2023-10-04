// https://www.codechef.com/problems/CBSPEED
package beginnerlevel
import "fmt"

func main() {
  
  var x, y int;
  fmt.Scan(&x, &y);
  if( y <= x ) {
    fmt.Println("NO");
  }else{
    fmt.Println("YES");
  }

}
