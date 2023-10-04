package beginnerlevel
// https://www.codechef.com/problems/CNTWRD
import "fmt"

func main() {
  
  var test int ;
  fmt.Scan(&test);
  var n, m int ;
  for test > 0 {
    fmt.Scan(&n, &m);
    fmt.Println(n * m );
    test --;
  }
}
