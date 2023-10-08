package beginnerlevel
//https://www.codechef.com/problems/HS08TEST
import "fmt"

func main() {

  var x int;
  var y float64;
  fmt.Scan(&x, &y);
  if(x % 5 == 0 && float64(x) + 0.5 <= y ){
    y = y - float64(x) - 0.5;
    fmt.Printf("%.2f", y);
  } else {
    fmt.Printf("%.2f", y);
  }


}
