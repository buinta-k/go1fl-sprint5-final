package personaldata

import "fmt"

type Personal struct {
	Name string
	Weight int
	Height int
}

func (p Personal) Print() {
	fmt.Printf("Имя: <%s>\n", p.Name)
	fmt.Printf("Вес: <%d>\n", p.Weight)
	fmt.Printf("Рост: <%d>\n", p.Height)
}
