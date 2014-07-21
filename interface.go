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

func (m *Man) myGender() string {
	return m.gender
}

func (m *Man) changeGender(){
	m.gender = "female"
}

type Woman struct {
	Person
}

func (w *Woman) myGender() string {
	return w.gender
}

func (w *Woman) changeGender() {
	w.gender = "male"
}

func (p *Person) myPointerType() string {
	return "human"
}

func PrintGender(h Human){
	println(h.myGender())
}

func changeGender(h Human) {
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
		changeGender(human_Arr[n])
		PrintGender(human_Arr[n])
		//println(human_Arr[n].gender)
	}

}


