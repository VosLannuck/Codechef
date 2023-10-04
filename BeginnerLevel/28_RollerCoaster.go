// https://www.codechef.com/problems/MINHEIGHT
package beginnerlevel

import "fmt"
func main() {
  
  var test int;
  var x, h int;

  fmt.Scan(&test);

  for test > 0 {
    fmt.Scan(&x, &h);
    if(x < h) {
      fmt.Println("NO");
    } else{
      fmt.Println("YES");
    }
    test --;
  }


}
