/*Задание 1: У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который определяет все ли числа в массиве data четные и записывает логический результат в переменную result.


*/

package main 

import( 
"fmt"
"strings"
"sort"
"strconv"
)

func main () {
//data_1:=[]int{2, 4, 6, 8, 10}
//data_2:=[]int{1,3,4,5}
//data_3:=[]int{4,56,78,2334,65}
//num_3:=10
//data_4:=[]int{150,2,34}
//num_4:=100
//data:=[]string{"Макс", "Дастин", "Майк", "Стив", "Билли"}
//data_6:=[]int{7,3,90,4,34}
data_7:=[]int{5,6,78,1,4,90}
//fmt.Println(even_or_odd(data_1))
//fmt.Println(odd_elements(data_2))
//fmt.Println(is_num_less_than_elinarray(data_3, num_3))
//fmt.Println(is_elinarray_less_than_num(data_4, num_4))
//fmt.Println(reversed_array(data))
//fmt.Println(sort_incr_arr(data_6))
fmt.Println(sort_reduce_arr(data_7))
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
/*Задание 5: У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных string.

Напишите код, который обращает порядок следования элементов массива data и записывает результат через запятую в переменную result.
*/
func reversed_array (data[]string) string {
    arr := []string{}
    for i:=0; i<len(data); i++ {
        arr=append(arr, data[len(data)-i-1])
    }
    result := strings.Join(arr, ", ")
    return result 
}
/*Задание 6: У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который сортирует числовые элементы массива data в порядке возрастания и записывает результат через запятую в переменную result.
*/
func sort_incr_arr (data_6[]int) string {
    var result string 
    stroki:=[]string{}
    arrSlice:=data_6[:]
    sort.Ints(arrSlice)
    for i:=0; i<len(arrSlice); i++ {
        stroki = append(stroki, strconv.Itoa(arrSlice[i]))
    }
    result = strings.Join(stroki, ", ")
    return result 

}
/*Задание 7:У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который сортирует числовые элементы массива data в порядке убывания и записывает результат через запятую в переменную result.
*/
func sort_reduce_arr (data_7 []int) string {
    finalArr:=[]string{}
    var result string 
    for j:=0; j<len(data_7); j++ {
        for i:=0; i<len(data_7)-1-j; i++ {
            if data_7[i] < data_7[i+1] { 
                m := data_7[i+1]; 
                data_7[i+1] = data_7[i];
                data_7[i] = m}
        }
    }
    for i:=0; i<len(data_7); i++ {
        finalArr = append(finalArr, strconv.Itoa(data_7[i]))
    }
    result = strings.Join(finalArr, ", ")
    return result 
}
/*Задание 8:
У вас есть переменная data, которая содержит входные пользовательские данные.

data - массив из элементов типа данных int.

Напишите код, который сортирует числовые элементы массива data в порядке возрастания, отсеивает дубликаты и записывает результат через запятую в переменную result.
*/

package main

import (
	"encoding/json"
	"fmt"
	//"strings"
    "bufio"
    "os"
    //"sort"
)

func main() {
    data := ReadInput()
    //var result string
    //arr := data
    //slice := []int{}
    for j:=0; j < len(data); j++ {
        for i:=0; i<len(data)-1-j; i++ {
            if data[i] > data[i+1] { m := data[i+1]; data[i+1] = data[i]; data[i] = m }
                //data[i], data[i+1] = data[i+1], data[i]}
        }
    }



    fmt.Println(removeDuplicates(data))
}
func removeDuplicates (arr[]int) []int {
    new_array:=[]int{}
        for i:=0; i<len(arr); i++ {
            if same_el(arr, i) == false { new_array = append(new_array, arr[i]) }
        }
    
  return new_array
}    
    
func same_el (arr[]int, number int) bool {
    for i:=0; i<len(arr); i++ {
        if arr[i] == number { return true }
    }
    return false 
}

func ReadInput() ([]int) {
    var data []int 
    var input string 
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input = scanner.Text()
    json.Unmarshal([]byte(input), &data)
    return data
}