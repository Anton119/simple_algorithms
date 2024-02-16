/*Задание 1: У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который определяет все ли числа в массиве data четные и записывает логический результат в переменную result.


*/

package main 

import "fmt"

func main () {
//data_1:=[]int{2, 4, 6, 8, 10}
//data_2:=[]int{1,3,4,5}
//data_3:=[]int{4,56,78,2334,65}
//num_3:=10
data_4:=[]int{150,2,34}
num_4:=100
//fmt.Println(even_or_odd(data_1))
//fmt.Println(odd_elements(data_2))
//fmt.Println(is_num_less_than_elinarray(data_3, num_3))
fmt.Println(is_elinarray_less_than_num(data_4, num_4))
}

func even_or_odd (data_1[]int) bool {
    var even int 
    var odd int 
    for i:=0; i<len(data_1); i++ {
        if data_1[i] %2 == 0 { even=data_1[i]
                           } else { odd=data_1[i] }
    }
    if odd %2 == 1 && even %2 == 0 { return false 
        } else { return true }

}
/*Задание 2: У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который определяет все ли числа в массиве data нечетные и записывает логический результат в переменную result.
*/
func odd_elements (data_2[]int) bool {
for i:=0; i<len(data_2)-1; i++ {
    if data_2[i] %2 !=0 && data_2[i+1] %2 !=0  { continue 
                                           } else { return false }
}
return true
}

/*Задание 3: У вас есть переменные num, data которые содержат входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который определяет все ли числа в массиве data больше num и записывает логический результат в переменную result.
*/
func is_num_less_than_elinarray (data_3[]int, num int) bool {
    for i:=0; i<len(data_3)-1; i++ {
        if data_3[i] > num && data_3[i+1] > num { continue
        } else { return false }
    }
    return true 
}
/*Задание 4: У вас есть переменные num, data которые содержат входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который определяет все ли числа в массиве data меньше num и записывает логический результат в переменную result.
*/
func is_elinarray_less_than_num (data_4[]int, num_4 int) bool {
    for i:=0; i<len(data_4); i++ {
        if data_4[i] >= num_4 {
            break
        } else { return true }
    }
    return false 
}

