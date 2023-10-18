package beginnerlevel
// https://www.codechef.com/problems/AVGPROBLEM
import "fmt"

func main(){

  var t int;
  var a, b , c float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b, &c );
    if( ( (a + b ) / 2) > c) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }

}
