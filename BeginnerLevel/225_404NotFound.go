package beginnerlevel

import ("fmt")
func main() {
  
 var x int ; 
  fmt.Scan(&x);
  if (x == 404) {
    fmt.Println("NOT FOUND");
  }  else {
    fmt.Println("FOUND");
  }
}
