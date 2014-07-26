package main

type Person struct {
	gender string
}

type Human interface {
	myGender() string
	changeGender()
}

func (p Person) myType() string {
	return "human"
}

type Man struct {
	Person
}

type Woman struct {
	Person
}

func (p *Person) myGender() string {
	return p.gender
}

func (p *Person) changeGender(){
	if p.gender == "male" { 
		p.gender = "female"} else
	if p.gender == "female" {
		p.gender = "male"}	
}

func PrintGender(h Human){
	println(h.myGender())
}

func alterGender(h Human) {
	h.changeGender()
}

func main() {
	m := new (Man)
	w := new (Woman)
	m.gender = "male"
	w.gender = "female"

	human_Arr := [...]Human{m,w}
	for n, _ := range (human_Arr)	{
		PrintGender(human_Arr[n])
		alterGender(human_Arr[n])
		PrintGender(human_Arr[n])
		//println(human_Arr[n].gender)
	}

}


