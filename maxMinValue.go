package main

import "fmt"

func main()  {

	var size int;
	fmt.Println("Enter the size of array");
	fmt.Scanln(&size);
	arr := make([]int, size)
	for i:=0;i<len(arr);i++{
		a:="please Enter the index";
		b:="element of array"
		message :=fmt.Sprintf("%s %d %s",a,i,b)	
		fmt.Println(message);
		fmt.Scanln(&arr[i])
	}

	max:=arr[0];
	min:=arr[0];

	for i:=1;i<len(arr);i++{
		if max<arr[i]{
			max=arr[i];
		}
		if min >arr[i]{
			min=arr[i]
		}
	}
	a:="max value is"
	b:="min value is"
	maximum:=fmt.Sprintf("%s %d",a,max);
	minmum:=fmt.Sprintf("%s %d",b,min);	
	fmt.Println(maximum);
	fmt.Println(minmum);


	
}


