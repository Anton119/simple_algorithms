/*Задание: У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который определяет все ли числа в массиве data четные и записывает логический результат в переменную result.


*/

package main 

import "fmt"

func main () {
data:=[]int{2, 4, 6, 8, 10}
var result bool
    var even int 
    var odd int 
    for i:=0; i<len(data); i++ {
        if data[i] %2 == 0 { even=data[i]
                           } else { odd=data[i] }
        if odd %2 == 1 && even %2 == 0 { result=false 
                       } else { result=true }
    }
	fmt.Println(result)

}