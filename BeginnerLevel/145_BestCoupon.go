package beginnerlevel
// https://www.codechef.com/problems/CHEAPFOOD
import "fmt"

func main() {
  
  var t,x int;
  fmt.Scan(&t);
  for t > 0 {
    
    fmt.Scan(&x);
    var TenPercentOff int =  ( x * 10 / 100 ) ;
    var MinusOneHundred int = 100;

    if(TenPercentOff > MinusOneHundred) {
      fmt.Println(TenPercentOff)
    } else  {
      fmt.Println(MinusOneHundred);
    }
    t--;
  }

}
