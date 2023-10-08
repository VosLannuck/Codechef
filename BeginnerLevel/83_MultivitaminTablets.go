package beginnerlevel

import "fmt"

func main() {

  var t,x,y int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y);
    if(x * 3 <= y) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }
}

