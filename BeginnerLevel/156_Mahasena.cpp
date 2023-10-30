#include<iostream>


int main() {
  
  int a, n ;
  std::cin>>n;
  int oddsTotal {0}, evenTotal{0} ;
  while(n--) {
    std::cin>>a;
    if(a % 2 == 0) {
      evenTotal +=1 ;
    }
    else {
      oddsTotal += 1;
    }
  }

  if(oddsTotal >= evenTotal) {
    std::cout<<"NOT READY\n";
  } else {
    std::cout<<"READY FOR BATTLE\n";
  }
  return 0;
}

