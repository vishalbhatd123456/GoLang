package main

import ("fmt")



func multiple(a,b string) (string,string){
	return a,b;
}


func main(){
	w1,w2 := "Hey","There";
	fmt.Println(multiple(w1,w2))
	
	
}
