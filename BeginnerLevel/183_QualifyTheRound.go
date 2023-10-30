package beginnerlevel
import "fmt"
func main() {
  
  var t, x ,a, b int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x,&a,&b);
    if(2* b + a >= x) {
      fmt.Println("Qualify");
    } else {
      fmt.Println("NotQualify")
    }
    
    t--;
  }
}

