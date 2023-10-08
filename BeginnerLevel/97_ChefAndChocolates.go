// https://www.codechef.com/problems/CCHOCOLATES
package beginnerlevel
import "fmt"
func main() {
  
  var t, x, y, z int;
  fmt.Scan(&t);

  for  t > 0 {
    fmt.Scan(&x, &y, &z)
    var totalCoins int = x * 5 + y * 10 ;
    fmt.Println( totalCoins / z );
    t--;

  } 

}



