// https://www.codechef.com/problems/MANGOES
package beginnerlevel


import "fmt"
func main() {
  var t, x, y, z int;
    fmt.Scan(&t);
  for t > 0 {
    
    fmt.Scan(&x,&y,&z);
    var afterTruck = z - y;
    var totalAble = afterTruck / x;

    fmt.Println(totalAble);
    t--;
  }


}  
