package beginnerLevel

// https://www.codechef.com/problems/TCKTFINE

import "fmt"

func main() {

  var t,x,q,p int;
  
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &p, &q);
    fmt.Println(x * ( p-q ));
    t--;
  }

}

