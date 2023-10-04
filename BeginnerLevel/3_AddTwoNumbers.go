package beginnerlevel;

import "fmt";

// Link : https://www.codechef.com/problems/FLOW001
func main() {

  var t int;
  var a, b int;
  fmt.Scan(&t);
  for i := 0 ; i < t ; i++ {
    fmt.Scan(&a, &b);
    fmt.Println(a + b);
  }; 
}

