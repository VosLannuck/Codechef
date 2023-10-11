// https://www.codechef.com/problems/PERFECTTRIO
package beginnerlevel
import "fmt"
func main() {
  var t,a,b,c int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a,&b,&c);
    if(a + b == c) {
      fmt.Println("YES");
    } else if (b + c == a) {
      fmt.Println("YES");

    } else if(c+a ==b ){
      fmt.Println("YES");
  

    }else {

      fmt.Println("NO");
    } 
    t--;
  }
}
