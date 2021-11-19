package accounts

import (
	"errors"
	"fmt"
)

//compile은 main에서 할꺼얌!

//Account struct
//private 이어서 변경해주자
//그 방법은 앞문자를 대문자로 변경하기
type Account struct {
	owner string 
	balance int 

}
//*return형이 어카운트의 복사본 
//NewAccount creates Account
func NewAccount(owner string) *Account	{
	  account := Account{owner: owner, balance: 0}
   
	  //실제로는 address 를 return하는 거에요!  복사본을 많이 만들면 어케된다? 느려진다.
	  return &account
}


//function과 method의 차이점 
//밑에는 method라고 불립니다.
//struct의 첫글자를 따서 변수명을 설정해야 한다는 규정이 있다.
//Deposit x amount on your account
//*Account를 복사하지말고 써라 이말입니다
func (a *Account) Deposit(amount int){

	a.balance += amount;

}


//Balance of your account
func (a Account) Balance() int{
return a.balance
}

//Withdraw  x amount form your 	account
func (a *Account) WithDraw(amount int)error{
	if a.balance < amount {
     return errors.New("Can't withdraw you are poor")
	}
    a.balance -= amount
	return nil
	//null 이나 none같은 거에요~
}



//Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner


}


//Owner of the account 
func (a Account) Owner() string{
	return a.owner
}


type Jun struct{
	name string
}

func NewJun(name string) *Jun{
	jun := Jun{name: name}

	return &jun
}

func(j Jun) String() string{
	return fmt.Sprint("Sprint 유형입니다!")
}

// func (a Account) String() string {
// 	// return fmt.Sprintln(a.Owner(), "s account. \nHas ", a.Balance())
// 	return "whatever you want"
// }