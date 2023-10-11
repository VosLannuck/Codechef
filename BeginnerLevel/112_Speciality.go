package beginnerlevel
//https://www.codechef.com/problems/SPECIALITY
import "fmt"

func main() {

  var t, x, y, z int; 
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &z);
    if(x > y && x > z) {
      fmt.Println("Setter");
    } else if (y > z && y > x) {
      fmt.Println("Tester");

    } else if(z > x && z > y) {

      fmt.Println("Editorialist");
  
    }
    t--;  
  } 



}
