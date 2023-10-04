// https://www.codechef.com/problems/KNGTOR
package beginnerlevel
import "fmt"
func main() {
  var test int;
  var n, m int;

  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&n, &m)
    fmt.Println(n * 5 + m * 7);
    test -- ;
  }

}
