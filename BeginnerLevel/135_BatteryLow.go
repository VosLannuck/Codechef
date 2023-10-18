
// https://www.codechef.com/problems/BATTERYLOW
package beginnerlevel

import "fmt"

func main () {

  var t, x int;
  fmt.Scan(&t);
  for t > 0  { 

    fmt.Scan(&x);
    if(x <=15 ) { 
      fmt.Println("Yes");
    }else{
      fmt.Println("No");
    }

    t--;

  }


}
