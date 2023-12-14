package beginnerlevel

import (
  "fmt"
)
func main () {

  var t int;
  fmt.Scan(&t);
  for t > 0 {
    var sx,sy, ex,ey int;
    fmt.Scan(&sx, &sy, &ex, &ey);
    if(sx != ex && sy != ey) {
      fmt.Println(1);
    } else {
      fmt.Println(2);
    }
    t--;
  } 
}


