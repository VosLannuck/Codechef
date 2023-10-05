package beginnerlevel
// https://www.codechef.com/problems/CHAIRS_
import "fmt"
import "math"
func main() {

  var test int;
  fmt.Scan(&test);

  for test > 0 {
    var x, y float64;
    fmt.Scan(&x, &y);
    if(x > y) {
      fmt.Println( (math.Abs(y-x)));
    }else {
      fmt.Println(0);
    }
    test--;
  }

}
