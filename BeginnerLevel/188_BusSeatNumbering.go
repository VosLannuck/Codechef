// https://www.codechef.com/problems/SEATNUMBER
package beginnerlevel

import "fmt"

func main() {

  var t, n int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n);

    if(n <= 15) {
      
      if(n <=10) {
        fmt.Println("Lower Double");
      } else {
      
        fmt.Println("Lower Single");
      }

    }else {
      
      if(n <= 25) {
      
        fmt.Println("Upper Double");
      } else {

        fmt.Println("Upper Single");
      }

    }


    t--;


  }
}
