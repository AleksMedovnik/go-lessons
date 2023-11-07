package lessons

import "fmt"

type superHero interface {
	superPower()
}

func useSuperPower(hero superHero) {
	hero.superPower()
}

type spiderMan struct{}

type batMan struct{}

func (s spiderMan) superPower() {
	fmt.Println("The spider jumps!")
}

func (b batMan) superPower() {
	fmt.Println("The bat is flying!")
}

// ########################################

type reader interface {
	read()
}

type writter interface {
	write(string)
}

func writeToStream(w writter, text string) {
	w.write(text)
}

func readFromStream(r reader) {
	r.read()
}

type file struct {
	text string
}

func (f *file) read() {
	fmt.Println(f.text)
}

func (f *file) write(message string) {
	f.text = message
	fmt.Println("Writing a string to a file", message)
}

// ############################

type readerWritter interface{
	reader
	writter
}

func reading(r readerWritter){
	r.read()
}

func Interfaces() {
	var bat superHero = batMan{}
	var spider superHero = spiderMan{}

	bat.superPower()    // The bat is flying!
	spider.superPower() // The spider jumps!

	useSuperPower(bat)    // The bat is flying!
	useSuperPower(spider) // The spider jumps!

	// ########################################

	f1 := &file{}
	writeToStream(f1, "Hello") // Writing a string to a file Hello
	readFromStream(f1)         // Hello

	reading(f1) // Hello

}
