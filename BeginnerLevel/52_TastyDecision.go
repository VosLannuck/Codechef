package beginnerlevel
// https://www.codechef.com/problems/TASTEDEC
import "fmt"

func main() {

  var t, x,y, choco, candies int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y);
    choco = x * 2;
    candies = y * 5;

    if(choco == candies) { 
      fmt.Println("Either")
    } else if(choco > candies) {
      fmt.Println("Chocolate")
    } else {
      fmt.Println("Candy")
    }
    t--;
  }

}
