// https://www.codechef.com/problems/RIGHTTHERE
package beginnerlevel
import "fmt"

func main() {
  
  var test int;

  var n , x int;

  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&n, &x);

    if(n <= x) {
      fmt.Println("Yes");
    }else {
      fmt.Println("NO");
    }

    test --;
  }

}
