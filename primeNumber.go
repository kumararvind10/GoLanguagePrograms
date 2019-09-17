package main

import "fmt";

func main()  {
	var num int;
	fmt.Println("Enter number to be checked prime or not");
	fmt.Scanln(&num);
	if num==1{
		fmt.Println("1 is not a prime Number");
		
	}else if num==2{
		fmt.Println("2  is a prime Number");
		
	}else{
		for i:=2;i<num;i++{
			if num%i==0{
				fmt.Println("is not a prime number");
				break;
				
			}else{
				fmt.Println("is a prime number");
				break;
				
			}
		}
	}

}