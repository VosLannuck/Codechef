package beginnerlevel
// https://www.codechef.com/problems/CMASKS
import "fmt"

func main() {

  var t, x, y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    x = x * 100;
    y = y * 10;

    if(x == y) {
      fmt.Println("Cloth");
    } else if( y < x) {
      fmt.Println("Cloth");
    } else {
      fmt.Println("Disposable");
    }

    t--;
  } 
}
