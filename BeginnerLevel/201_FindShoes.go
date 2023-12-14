package beginnerlevel

import ("fmt")
// https://www.codechef.com/problems/FINDSHOES
func main () {

  var t int;
  var n , m int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &m);
    if(m == 0) {
      fmt.Println(n * 2);
    } else if(n > m) {
      fmt.Println(n + (n-m) );
    } else {
      fmt.Println(n);
    }

    t--;
  }

}
