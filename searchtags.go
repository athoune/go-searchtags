package main

var searchQueue chan queryAnswer

func main() {
	searchQueue = make(chan queryAnswer, 40)
	go StartSearch()
	go StartSearch()
	go StartSearch()
	go StartSearch()
	loadData()
	startHttp()
}
