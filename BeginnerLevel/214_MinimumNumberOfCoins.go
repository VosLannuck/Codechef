package beginnerlevel
// https://www.codechef.com/problems/MINCOINS
import "fmt"

func main() {

  var t, x int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    if( x % 10 == 0  ) {
      fmt.Println( x / 10);
    
    } else if (x % 5 == 0) {
      fmt.Println( (x / 10 ) + 1);
    } else {
      fmt.Println("-1")
    } 
    t--;
  }
  
}
