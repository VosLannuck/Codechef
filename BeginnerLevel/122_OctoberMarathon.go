package beginnerlevel
// https://www.codechef.com/problems/OCTATHON
import "fmt"

func main() {

  var x int ;
  fmt.Scan(&x);
  if(x < 3) {
    fmt.Println("GOlD");
  }else if(x>=3 && x < 6)  {
    fmt.Println("SILVER");
  } else{
    fmt.Println("BRONZE");

  } 

}
