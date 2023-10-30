// https://www.codechef.com/problems/DISCUS

package beginnerlevel
import "fmt"
func main () {
  
  var t,a,b,c int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b, &c);
    max := a;
    if(max  <b) {
      max = b;
    } 

    if(max < c) {
      max = c;
    }
  
    fmt.Println(max);


    t--;
  }




}
