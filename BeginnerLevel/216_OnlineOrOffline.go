// https://www.codechef.com/problems/FOODPLAN


package beginnerlevel

import "fmt"
func main() {
  
  var t int 

  var n ,m float64;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n, &m);
    var finalOnlineCosts float64 = n - ( n * 10 / 100);
    if(finalOnlineCosts < m) {
      fmt.Println("ONLINE");
    } else if (finalOnlineCosts > m ) {
      fmt.Println("DINING");
    } else {
      fmt.Println("EITHER");
    }
    t--;
  }

}
