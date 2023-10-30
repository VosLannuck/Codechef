package beginnerlevel
// https://www.codechef.com/problems/NIBBLE
import "fmt"

func main() {

  var t, x int;
  fmt.Scan(&t);

  for t > 0 {

    fmt.Scan(&x);
    if(x % 4 == 0 ) {
      fmt.Println("GOOD");
    } else {
      fmt.Println("NOt Good");
    }
    t--;
  }

}
