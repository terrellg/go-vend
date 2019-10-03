package main

type machine struct {
	model    string
	address  address
	capacity int
}

func (m machine) String() string {
	return "This machine is located at \n" + m.address.String()
}
