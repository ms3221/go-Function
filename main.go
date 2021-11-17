package main

import (
	"fmt"
	"strings"
)

//istrue := true  함수밖에서는 될수가 없어요!!! 축약형은 오로지 함수안에서만 가능합니다
//그래서 밖에서는 이렇게 작성을 해야합니다.

//자바는 int multiply(int a, int b) 이렇게 반환해주는 타입과 들어오는 타입을 지정해주는데 go는 정반대로 한다!! 주의!!
 
func multiply(a, b int) int{
	return a * b
}


//go의 가장 멋지고 cool한 부분은 많은 return 값을 반환한다는 것이다.
// func lenAndUpper(name string)(int,string){
// 	return len(name), strings.ToLower(name);
// 	}

func lenAndUpper(name string)(length int, uppercase string){
	//defer이라는 기능은 func이 끝났을 때 추가적으로 무엇인가 동작 할 수 있도록 할 수 있다. 
	defer fmt.Println("I'm done") //function끝나고 나서 코드 실행!
    length = len(name);    //위에처럼 Length, upperacse를 작성해주면 저것을 return할것이라고 명시되어있는 것과 같다. naked return
	uppercase = strings.ToUpper(name)   //naked return 내가 직점 무엇을 리턴할 것인지를 명시하지 않아도 됩니다.
	return length, uppercase;
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// func main() {
	
// 	// // const name string = "nico"

// // 	//name := "123" // === var name string = "nico"

// //     totalLenght, upperName := lenAndUpper("sia")
// // 	fmt.Println(totalLenght,upperName);
   
// // 	// repeatMe("nico","lynn","dal","marl","flynn")

// }



//============================= roop =================================

// //for에 대해서 공부해 보자구요! 
// func superAdd(numbers ...int) int {
//     //roop을 만드는 방법.
// 	// for index, number := range numbers{
// 	// 	fmt.Println(index,number)
// 	// }

// 	for i := 0; i<len(numbers); i++{
// 		fmt.Println(numbers[i])
// 	}

// 	total := 0;
// 	//_이부분은 index를 ignore하는 부분입니다.
// 	// for _, number := range numbers{
// 	// 	total += number
// 	// }

// 	return total
// }

// func main(){
// result := superAdd(1,2,3,4,5,6);

// fmt.Println(result)
// }


//============================= roop end-- =================================


//============================= if/else-- =================================

// func canIDrink(age int) bool {
//   if koeranAge := age + 2; koeranAge < 18 {
// 	  return false
//   }
// 	  return true
  
// }

// func main(){
// 	fmt.Println(canIDrink(15))
// }
  
//============================= if/else end-- =================================


//============================= switch -- =================================

// func canIDrink(age int) bool {

// 	switch koreanAge := age+2; koreanAge {
// 	case 18 : 
// 	   return false
//     case 20 : 
// 	   return true
// 	case 50 :
// 		return false
// 	}

// 	return false
	
//   }

  
//   func main(){
// 	  fmt.Println(canIDrink(17))
//   }

//============================= end -- =================================



//============================= pointer =================================

//   func main()  {

// 	//이렇게 &하면 주소값을 찾을 수가 있어용1
// 	// fmt.println(b) == 0x12341243 주소값이 생성된다. 이값을 내가 볼 수 있는 정수값으로 반환할려면 *을 붙여줍시다
// 	// a의값을 업데이트해주면 b는 a의 주소값을 참조하기 때문에 같이 변합니다. 
// 	// 그리고 또한 가지방법은 b의 값을 업데이트해서 a의 값을 바꾸는 건데 그건 *b = 12 이런식으로 업데이트 해주면 됩니다.
// 	  a :=2
// 	  b := &a
//       *b = 20;
	 
	  

// 	  fmt.Println(a,*b)
//   }



	//============================= end -- =================================




//============================= array =================================

	// func main(){
	// 	names := []string{"nico","lynn","dal"}
	// 	names = append(names,"flyn");   // append는 새로운 값이 추가된 slice를 return 합니다. go는 배열에서 	push, pop이라는 함수를 사용하지 않습니다.
	// 	//append()라는 함수를 사용합니다.
	// 	fmt.Println(names)   

	// 	for _,name := range names {
    //       fmt.Println(name)
	// 	}
	// }


//============================= end -- =================================


//============================= map =================================

// func main(){
//   jun := map[string]string{"name":"nico","age":"12"}

//   for _,value := range jun{
// 	  fmt.Println(value);
//   }
// }

//============================= end -- =================================

//============================= struct =================================

type person struct{
	name string
	age int
	favFood []string
}

func main(){
	favFood := []string{"kimchi","ramen"}
	jun := person{name: "jun",age: 18,favFood:favFood}

	fmt.Println(jun)
}

//============================= end -- =================================







