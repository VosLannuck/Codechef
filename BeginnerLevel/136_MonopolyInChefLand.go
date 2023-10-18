package beginnerlevel

import "fmt"

func main() {
 
  var t, r1,r2,r3 int;
  fmt.Scan(&t);

  for t > 0 {
  
    fmt.Scan(&r1,&r2,&r3);
    if( r1 > r2  + r3 ) { 
      
      fmt.Println("YES");
    } else if( r2 > r1 + r3 ) { 

      fmt.Println("YES");
    } else if ( r3 > r1 + r2 ) {

      fmt.Println("YES");

    }else {


      fmt.Println("NO");
    }



  } 
}
