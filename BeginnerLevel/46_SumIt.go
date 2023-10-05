// https://www.codechef.com/problems/SUMM
package beginnerlevel
import "fmt"
func main() {
  
  var test, a,b,c int;

  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&a , &b, &c);
    if(a + b == c) {
      fmt.Println("YES");
    }else {
      fmt.Println("NO");
    }
    test --;
  }

}
