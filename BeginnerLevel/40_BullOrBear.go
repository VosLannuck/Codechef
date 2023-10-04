package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/BULLBEAR
func main() {
  
  var test int;
  var x, y int;
  fmt.Scan(&test);
  for test > 0 {

    fmt.Scan(&x,&y);
    if(y - x ==  0 ) {
      fmt.Println("NEUTRAL");
    } else if (y - x > 0) {
      fmt.Println("PROFIT");
    } else  {
      fmt.Println("LOSS");
    }
    test --;
  }

}
