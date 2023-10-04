package beginnerlevel
// https://www.codechef.com/problems/PAR2
import "fmt"
func main() {
  
  var test int;

  var n int;
  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&n);
    if(n %2 == 0 ) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }

    test --;
  }

}
