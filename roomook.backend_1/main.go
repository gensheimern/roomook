package main

	func main() {
	router := GinRouter().SetupRouter()

	go hub.run()
	router.Run(":3000")
		
}
	




