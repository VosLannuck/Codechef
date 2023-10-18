package beginnerlevel

import "fmt"

func main() {

  var t,w,x,y,z int; 

  fmt.Scan(&t);

  for t > 0 {


    fmt.Scan(&w,&x,&y,&z);
    var totalWater = y * z  + w;
    if(totalWater > x ) {
      fmt.Println("overFlow");
    }else if (totalWater < x) {
      fmt.Println("Unfilled");
    } else {
      fmt.Println("filled");
    }
    
    t--;
  }
  
}
