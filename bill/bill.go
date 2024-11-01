package bill


type bill struct {
	name string
	items map[string]float64
	tip float64
}

func NewBill(name string) bill{
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0, // map[string]float64{}
	}
return b
}