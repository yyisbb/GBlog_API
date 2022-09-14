package cmd

func Run(configCenter func()) {
	go configCenter()
	select {}
}
