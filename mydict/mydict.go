package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

type Money int 



//Search
func(d Dictionary) Search(word string)(string,error){
  value, exits := d[word] 
  if  exits {
	  return value, nil
  }else{
      return "",errNotFound
  }
}