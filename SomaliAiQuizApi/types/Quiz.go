package types

type Quiz struct {
	ResponseCode int      `json:"response_code"`
	Results      []Result `json:"results"`
}

type Result struct {
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Category         string   `json:"category"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

func NewQuiz() *Quiz {
	return &Quiz{}
}

// response_code	0
// results
// 0
// type	"multiple"
// difficulty	"hard"
// category	"Science &amp; Nature"
// question	"Which of the following is the term for &quot;surgical complications resulting from surgical sponges left inside the patient&#039;s body?"
// correct_answer	"Gossypiboma"
// incorrect_answers
// 0	"Gongoozler"
// 1	"Jentacular"
// 2	"Meupareunia"
// 1
// type	"boolean"
// difficulty	"medium"
// category	"History"
// question	"The Hundred Years&#039; War was fought for more than a hundred years."
// correct_answer	"True"
// incorrect_answers
// 0	"False"
// 2
// type	"multiple"
// difficulty	"medium"
// category	"Entertainment: Japanese Anime &amp; Manga"
// question	"In &quot;Highschool DxD&quot;, what is the name of the item some humans are born with?"
// correct_answer	"Sacred Gear"
// incorrect_answers
// 0	"Imperial Arm"
// 1	"Hallowed Relic"
// 2	"Blessed Artifact"
// 3
// type	"multiple"
// difficulty	"medium"
// category	"Entertainment: Film"
// question	"Who voiced Metalbeard in &quot;The Lego Movie&quot;?"
// correct_answer	"Nick Offerman"
// incorrect_answers
// 0	"Liam Neeson"
// 1	"Morgan Freeman"
// 2	"Will Arnet"
// 4
// type	"multiple"
// difficulty	"medium"
// category	"Entertainment: Television"
// question	"Which boxing personality was one of the presenters in the 1999 revival of It&#039;s a Knockout?"
// correct_answer	"Frank Bruno"
// incorrect_answers
// 0	"Henry Cooper"
// 1	"Muhammad Ali"
// 2	"Joe Fraiser"
// 5
// type	"multiple"
// difficulty	"medium"
// category	"Entertainment: Board Games"
// question	"Who was world chess champion between 1894 and 1921"
// correct_answer	"Emanuel Lasker"
// incorrect_answers
// 0	"Jos&eacute; Ra&uacute;l Capablanca"
// 1	"Wilhelm Steinitz"
// 2	"Bobby Fischer"
// 6
// type	"multiple"
// difficulty	"hard"
// category	"Entertainment: Cartoon &amp; Animations"
// question	"Who played Stan&#039;s dog in the South Park episode &quot;Big Gay Al&#039;s Big Gay Boat Ride&quot;?"
// correct_answer	"George Clooney"
// incorrect_answers
// 0	"Jay Leno"
// 1	"Matt Stone"
// 2	"Robert Smith"
// 7
// type	"multiple"
// difficulty	"medium"
// category	"Geography"
// question	"Gibraltar, located just south of the Iberian peninsula, is a territory of which West Europe country?"
// correct_answer	"United Kingdom"
// incorrect_answers
// 0	"Spain"
// 1	"Portugal"
// 2	"France"
// 8
// type	"multiple"
// difficulty	"medium"
// category	"Entertainment: Music"
// question	"The &#039;In the Flesh&#039; tour was used in support of what Pink Floyd album?"
// correct_answer	"Animals"
// incorrect_answers
// 0	"The Wall"
// 1	"Wish You Were Here"
// 2	"The Final Cut"
// 9
// type	"boolean"
// difficulty	"hard"
// category	"Science: Computers"
// question	"The IBM PC used an Intel 8008 microprocessor clocked at 4.77 MHz and 8 kilobytes of memory."
// correct_answer	"False"
