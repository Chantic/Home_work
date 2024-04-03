package people

var Vasya string = "Vasya"
var Shantal string = "Shantal"
var Victory string = "Victory"
var Karin string = "Karin"
var Erik string = "Erik"
var Milana string = "Milana"
var Aziz string = "Aziz"

// Box := []string{Vasya, Shantal, Victory, Karin, Erik, Milana, Aziz}
type People []string

//1 создать переменную с типом People и запихать туда все имена

var PeopleList People = []string{Vasya, Shantal, Victory, Karin, Erik, Milana, Aziz}

func (p *People) Sort() {
	//var Box2 []PeopleList= []PeopleList{Vasya, Shantal, Victory, Karin, Erik, Milana, Aziz}
	for i := 0; i < len(*p)-1; i++ {
		for j := 0; j < len(*p)-1-i; j++ {
			if []rune((*p)[j])[0] > []rune((*p)[j+1])[0] {
				(*p)[j], (*p)[j+1] = (*p)[j+1], (*p)[j]
			}

		}
	}
}

func (p *People) Let() string {
	first := (*p)[0]
	return first
}
