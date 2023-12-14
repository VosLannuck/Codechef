package beginnerlevel
// https://www.codechef.com/problems/CHEFEREN

import "fmt"


func main() {

  var t,n,a,b int;

  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n, &a, &b);
    var total int = 0;
    for n > 0 {
      if(n % 2 == 0 ) {
        total += a;
      }else {
        total += b; 
      }
      n--;
    } 
    fmt.Println(total);
    t--;
  }








}
