package beginnerlevel

import (
	"fmt"
)

// https://www.codechef.com/problems/CHEFAPPS
func main () {
  var t int;
  var s, x, y, z float64; 
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&s, &x, &y, &z );
    storageLeft := s - ( x + y);
    if(storageLeft >= z) {
      fmt.Println(0);
    }else {
      storageLeft_1 := storageLeft  + x;
      storageLeft_2 := storageLeft + y;
      if (storageLeft_1 >= z || storageLeft_2 >= z ) {
        fmt.Println(1);
      } else {
        fmt.Println(2);
      }

    }
    t--;
  }
}
