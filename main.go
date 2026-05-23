package main
func main() {
	args := os.Args[1:]

	run_hour := args[0]

	log.Printf("Hour passed: %s\n", run_hour)

	// Regexes
	re_hour, _ := regexp.Compile(`^(\d{2}):(\d{2})$`)
	re_special, _ := regexp.Compile(`^([noteaimdgh]+)$`)
	re_nplus, _ := regexp.Compile(`^(now)\s\+\s(\d+)\s([minuteshordaywk]+)$`)
	re_tomorrow, _ := regexp.Compile(`^(\d{2}):(\d{2})\stomorrow$`)
	re_day, _ := regexp.Compile(`^(\d{2}):(\d{2})\s([mondaytueswhrfi]+)$`)

}
