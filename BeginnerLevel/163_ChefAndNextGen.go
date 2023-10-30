package beginnerlevel

import "fmt"

func main() {

  var t, a , b , x , y int;
  fmt.Scan(&t);
  
  for t > 0 {
    fmt.Scan(&a, &b, &x, &y);
    var totalNeeded int = a * b;
    var exists int = x * y;
    if(totalNeeded > exists) {
      fmt.Println("NO");
    } else{
      fmt.Println("YES");
  }
    t--;
  }

} 
