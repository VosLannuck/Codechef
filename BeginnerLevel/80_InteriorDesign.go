package beginnerlevel
// https://www.codechef.com/problems/INTRDSGN
import "fmt"

func main() {

  var x1,y1,x2,y2 int;
  var t int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x1, &y1, &x2, &y2);
    if(x1 + y1 > x2 +y2) {
      fmt.Println( x2 + y2);
    } else {
      fmt.Println( x1 + y1);

    }
    t--;
  }

}
