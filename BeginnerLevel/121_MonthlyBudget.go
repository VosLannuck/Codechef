package beginnerlevel
// https://www.codechef.com/problems/BUDGET_

import "fmt"

func main() {

  var t, x , y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x,&y);
    y = y * 30 ;
    if(y <= x) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }

}
