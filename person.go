package person

type Person struct {
	Id int

	Name string

	Hobby string
}

func (p Person) GetId() int {

	return p.Id

}

func (p *Person) SetId(id int) {

	p.Id = id

}

func (p Person) GetName() string {

	return p.Name

}

func (p *Person) SetName(name string) {

	p.Name = name

}

func (p Person) GetHobby() string {

	return p.Hobby

}

func (p *Person) SetHobby(hobby string) {

	p.Hobby = hobby

}

func (p Person) ToString() string {

	return "Halo nama saya" + p.Name + ", hobi saya adalah " + p.Hobby

}
func GetPerson() Person {
	return Person{}
}
