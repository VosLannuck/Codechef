package beginnerlevel
 // https://www.codechef.com/problems/CHEFGAMES
import "fmt"


func main(){
  
  var t, r1, r2, r3, r4 int;
  fmt.Scan(&t);
  for t > 0 { 
    fmt.Scan(&r1, &r2, &r3, &r4);

    if(r1 == 0 && r2 == 0 && r3 == 0 && r4 == 0 ) {
      fmt.Println("IN");
    } else {
    
      fmt.Println("OUT");
    }
    t--;
  }

}
