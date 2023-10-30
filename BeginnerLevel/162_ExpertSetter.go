package beginnerlevel
// https://www.codechef.com/problems/EXPERT
import "fmt"

func main() {

  var t int 
  var x , y float64;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y);
    var approval float64 = (y / x) * 100;
    if(approval >= 50) {
      fmt.Println("YES");
    }  else {
      fmt.Println("NO");
    }
    t--;  
  }
    
}
