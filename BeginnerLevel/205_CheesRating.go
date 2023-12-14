package beginnerlevel
// https://www.codechef.com/problems/C_RATING
import "fmt"
func main () {
  
  var  t, x , y int64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    ans := y- x ;
    if( ans % 8 == 0 ) {
      fmt.Println(ans/8);
    } else if (ans % 8 != 0 ) {
      fmt.Println( (ans / 8) + 1);
    } else if (ans < 8 ) {
      fmt.Println(1);
    }
    t--;
  }


}
