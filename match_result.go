/*Вы создаете программу для анализа результатов спортивных матчей и подсчета очков заданной команды.
Результаты матчей хранятся в массиве results.
Каждый матч имеет один из следующих результатов:
"w" - выиграл
"l" - проиграл
"d" - ничья

Победа добавляет три очка, ничья - одно очко, а проигранный матч не добавляет очков.

Ваша программа должна принять на вход результат последнего матча и добавить его в массив результатов. 
После этого необходимо вычислить и вывести общее количество очков, 
набранных командой по результатам матчей из массива results.
*/
package main

import (
"fmt"
// "strconv"
)

func main() {
    //var total_points int
    //var last_match int
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    //fmt.Scan(&last_match)
    //results = append(results, strconv.Itoa(last_match))
    fmt.Println(show_points(results))
    

}


"43a46" -> 4346

func show_points(points [] string) int {
    total_points:=0
    for i:=0; i<len(points); i++ {
//    a, _ := strconv.Atoi(points[i])
    a:=666;
    switch {
        case points[i] == "w": a = 3 
        case points[i] == "l": a = 0
        case points[i] == "d": a = 1
    }
    total_points += a
    }
    return total_points
}