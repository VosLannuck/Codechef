
// https://www.codechef.com/problems/JENGA
package beginnerlevel

import "fmt"
func main() {


  var t , n , x int;
  fmt.Scan(&t); 

  for t > 0 {
    fmt.Scan(&n, &x);
    if(n > x) {
      fmt.Println("NO");
    }else {
      if(x % n == 0) {
        fmt.Println("YES"); 
      } else {
        fmt.Println("NO");
      }
    }
    t--;
  }



}
