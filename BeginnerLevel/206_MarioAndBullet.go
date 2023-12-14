package beginnerlevel
// https://www.codechef.com/problems/BULLET
import "fmt"
func main () {

  var t , x , y ,z int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &z);  
    
    timeTheBulletHits := y / x;
    if(timeTheBulletHits > z) {
      fmt.Println(0); 
    } else {
      fmt.Println(z - timeTheBulletHits);
    }
    



    t--;
  }

}
