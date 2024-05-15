package main

	func main() {
		cards := newDeckFromFile("savedFile")
		
		cards.shuffle()
		cards.print()
}


