// https://www.codechef.com/problems/PRACTICEPERF
package beginnerlevel

import "fmt"

func main() {

  var p1, p2 ,p3, p4, counter int ;
  fmt.Scan(&p1, &p2, &p3, &p4);

  if(p1 >= 10 ) { counter ++; }
  if(p2 >= 10 ) { counter++; }
  if(p3 >= 10 ) {counter ++;}
  if(p4 >= 10 ) { counter ++;}
  fmt.Println(counter);

}

