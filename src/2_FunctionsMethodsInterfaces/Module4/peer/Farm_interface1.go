package main
import "fmt"
import "bufio"
import "os"
import "strings"

type Animal interface{
	//init(string)
	Eat()
	Move()
	Speak()
}

type animals  string
type animals0 string
type animals1 string



func (cow *animals) Eat(){
	
	fmt.Println("grass")

}
func (cow *animals) Speak(){
	
	fmt.Println("moo")
    
}
func (cow *animals) Move(){
	
	fmt.Println("walk")
    
}


func (bird *animals0) Eat(){
	
	fmt.Println("worms")

}
func (bird *animals0) Speak(){
	
	fmt.Println("peep")
    
}
func (bird *animals0) Move(){
	
	fmt.Println("fly")
    
}
func (snake *animals1) Eat(){
	fmt.Println("mice")
}
func (snake *animals1) Speak(){
	fmt.Println("hsss")
}
func (snake *animals1) Move(){
	fmt.Println("slither")
}

func print_0(str string, a Animal){
	if str=="speak"{
		a.Speak()
	}
	if str=="move"{
		a.Move()
	}
	if str=="eat"{
		a.Eat()
	}


}

func inslice(a *[]string, b string) bool{
	for _,v :=range *a{
		if v==b{
			return false
			} }
			return true
}

func info_slice(sl1 *[]string, name string, animal0 string, animal1 string){
	if animal1==animal0{
		if  inslice(sl1, name){	
        
			fmt.Println("Created it!")
			(*sl1)=append((*sl1),name)
		}
}}

func main(){


//var cow_ptr animals
var inter0 Animal
var cow *animals
var bird *animals0
var snake *animals1
//var bird animals0
//var snake animals1


//inter0=cow
//inter0.Eat()

//(*inter0_ptr).Eat()

sl_cow:=make([]string,0,3)
sl_bird:=make([]string,0,3)
sl_snake:=make([]string,0,3)
fmt.Println("!!!Be aware of entering the commands as required. See examples below.!!!")
fmt.Println("Ends the program:> stop")
fmt.Println("Example create: >newanimal Jack bird")
fmt.Println("Example query:  >query Jack eat")


for{
	fmt.Printf(">")
    scanner:=bufio.NewScanner(os.Stdin)
    scanner.Scan()
	res:=strings.Split(scanner.Text()," ")
	if res[0]=="stop"{
		break
	}
	if res[0]=="newanimal"{
		info_slice(&sl_cow,res[1],res[2],"cow")
		info_slice(&sl_bird,res[1],res[2],"bird")
		info_slice(&sl_snake,res[1],res[2],"snake")
	}

	  
	if res[0]=="query"{
		if !(inslice(&sl_cow,res[1])){
			inter0=cow
			print_0(res[2],inter0)
		}
		if !(inslice(&sl_bird,res[1])){
			inter0=bird
			print_0(res[2],inter0)
		}
		if !(inslice(&sl_snake,res[1])){
			inter0=snake
			print_0(res[2],inter0)
		}
      
	}  

}
}



