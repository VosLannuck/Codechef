package beginnerlevel
// https://www.codechef.com/problems/INSURANCE
import "fmt"

func main(){

  var t, x, y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if( y <= x) {

      fmt.Println(y);
    }else {
      fmt.Println(x);
    }
    t--;
  }
} 
