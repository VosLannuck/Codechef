package beginnerlevel
// https://www.codechef.com/problems/AIRLINES
import "fmt"

func main() {
  var t, x, y, z int;
  fmt.Scan(&t);

  for t > 0 {

    fmt.Scan(&x, &y, &z );
    var total int =  x * 10;
    if( y > x) {
      fmt.Println(total * z); 
    } else { 

      fmt.Println(( y) * z);

    }

    t--;

  } 
}
