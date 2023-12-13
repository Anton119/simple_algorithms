package main 

import "fmt"

func main (){
a:= []int  {5,2,3,4}
mass1:= [] int {1,3,5,6}
mass2:= [] int {4,6,9}
a[0]=42
//arr:= [] int {5,2,7,8,4,1}
fmt.Println(mass1)
fmt.Println(mass2)
fmt.Println(merge(mass1, mass2))  // res -- фактический аргумент (actual parameter)
}
func mas(b [] int) bool {
for i:=0; i<len(b)-1; i++ {
if b[i]<b[i+1] {continue
} else {return false}
}
return true 
}

//0 1 2 3 4 8

func sort (arr [] int) [] int {
for mas(arr) != true {
for i:=0 ; i<len(arr)-1 ; i++ {
if arr[i] >  arr[i+1]  { 
x:=arr[i]
arr[i]=arr[i+1]
arr[i+1]=x
}
} 
}
return arr
}

func max (arr [] int) int {
   m:=arr[0]
   for i:=0; i<len(arr); i++ {
      if m < arr[i] { m=arr[i]
      } else { continue }
   }
   return m
}

func selection_sort (arr [] int) [] int {
   for j:=0; j < len(arr); j++ { 
      m:=arr[j]
      ind:=j
      for i:=j; i<len(arr); i++ {
           if m < arr[i] { m=arr[i]; ind=i 
           } else { continue }
           }
      tmp:=arr[ind]
      arr[ind]=arr[j]
      arr[j]=tmp
      }
    return arr
 }
          
          


func quick_sort(arr [] int) [] int {
if len(arr) == 0 {return arr}  
less:= [] int {}
more:= [] int {}
for i:=1; i<len(arr); i++{
  if arr[0]>arr[i] { less = append(less, arr[i]) 
  } else { more =  append(more, arr[i]) } 
}
less = quick_sort(less)
more = quick_sort(more)
arr =  append(append(less, arr[0]),  more...)

return arr 
}



func merge(mass1 [] int, mass2 [] int) [] int {  // формальные аргументы (formal arguments (parameters))
res:= [] int {}
for len(mass1)!= 0 && len(mass2) != 0  {
if mass1[0] < mass2[0] { res = append(res, mass1[0])
mass1 =  mass1[1:] } else { res = append(res, mass2[0])
mass2 =  mass2[1:] }
}
if len(mass1) == 0 { res = append(res, mass2...)
} else if len(mass2) == 0 {res = append(res, mass1...)
}
return res
}


/*

 5 8 9 

 4 7

0 2

5 9 8 2 1 4 9

5 9 8 2 
1 4 9

 

1 2 4 5 8 9 9


a=y*z
x=a+2


x=(a=y*z)+2

0 1 2 5 6 7

0 2 3 5 9


5 0 2 7 3 9 8 4

0 2 3 4

7 8 9  

0 2 3 4 5 7 8 9



*/ 




/* a=23
b=45
c = a

c = 23 
a = 23
b = 45 

a=b
x=45
i=45
i+1=23

i=i+1
x=45
i=23
i+1=23*/

//fmt.Println(a[2]


// asm
// C++ -- приборы/игры

// Go -- промежуточное
// PHP, JS, Python -- скрипты

// Go -> JS (PHP, PYTHON) -> C -> c++ -> asm -> clojure, scala