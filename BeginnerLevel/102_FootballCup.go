package beginnerlevel
//https://www.codechef.com/problems/FOOTCUP

import "fmt"

func main() {

  var t, x, y int ;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if(x > 0 && y > 0 && x == y ) {
      fmt.Println("Yes");
    }else {
      fmt.Println("No");
    }
    t--;
  }

}
