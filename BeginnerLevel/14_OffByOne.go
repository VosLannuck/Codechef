package beginnerlevel

import "fmt"
import "strconv"
//https://www.codechef.com/problems/OFFBY1 
func main() {
  var a , b  int ;
  fmt.Scan(&a, &b);
  fmt.Printf("%s", strconv.Itoa(a+b) + "1");

}
