// https://www.codechef.com/problems/CANCHEF
package beginnerlevel
import "fmt"

func main (){
  var t , x  , y int ;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x,&y);
    var total = 2 * y;
    var totalPetrol = 15 * x;

    if(totalPetrol >= total){
      fmt.Println("yes");

    }else{

      fmt.Println("no");
    } 
    t--;
  }
  


}
