package beginnerlevel

import "fmt"

func main() {

  var t, a, b, c int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a,&b,&c);
    if(a <= b && c <=b ) {
      fmt.Println("YES");
    } else{
      fmt.Println("No");
    }
    t--;
  }


}
