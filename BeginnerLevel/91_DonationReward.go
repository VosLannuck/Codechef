package beginnerlevel
//https://www.codechef.com/problems/DOREWARD 
import  "fmt"

func main() {
  var t, x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    if(x <= 3 ) {
      fmt.Println("BRONZE");
    } else if (x > 3 && x<=6) {
      fmt.Println("SILVER");
    
    } else {
      fmt.Println("GOLD");
    }
    t--;
  }  
}
