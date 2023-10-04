package beginnerlevel
// https://www.codechef.com/problems/PRIZEPOOL
import "fmt"

func main() {
  var test int;

  var y, x int;
  const TOP_PARTICIPANTS int = 10
  const OTHER_PARTICIPANTS int = 90
  fmt.Scan(&test);

  for test > 0 {
    fmt.Scan(&x, &y);
    fmt.Println(y * OTHER_PARTICIPANTS + x * TOP_PARTICIPANTS);
    test --;
  }
}
