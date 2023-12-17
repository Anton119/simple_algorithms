//задан числовой массив. сосчитать и вернуть, сколько различных чисел в этом массиве. 
//Например, в массиве 5, 7, 5 различных чисел два (5 и 7).
package main 

import "fmt"

func main () {
	arr:= [] int {5,6,1,5,7}
	dif_number:=0
	n:=arr[0]
	for i:=0; i<len(arr); i++ {
		if arr[0] != arr[i] { n = arr[i] 
			dif_number = i
		} else { continue }
		fmt.Println(n)
		fmt.Println(dif_number)
	}
}