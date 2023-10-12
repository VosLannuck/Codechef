package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/FLOW004
func getFirst(value int) int {
  var returnedVal int = 0 ;
  for value > 0 {
    returnedVal = value % 10;
    value = value / 10;
  }

  return returnedVal
}

func getLast(value int) int {
  return value % 10;
}
func main() {
  var t, n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n);
    var first int = getFirst(n);
    var end int = getLast(n);
    fmt.Println(first + end );
    t--;
  }   


}
