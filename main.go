package main
func main() {
	args := os.Args[1:]

	run_hour := args[0]

	log.Printf("Hour passed: %s\n", run_hour)

}
