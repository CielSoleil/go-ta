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

	switch {
	case re_hour.MatchString(run_hour):
		matches := re_hour.FindStringSubmatch(run_hour)
		hours, minutes := matches[1], matches[2]

		// Convert hours to int for special evaluation
		cHours, err := strconv.Atoi(hours)

		if err != nil {
			log.Fatal("Couldn't convert to int")
		}

		switch cHours {
		case 24:
			hours = "00"
		case 25:
			hours = "01"
		case 26:
			hours = "02"
		case 27:
			hours = "03"
		case 28:
			hours = "04"
		case 29:
			hours = "05"
		case 30:
			hours = "06"
		}

	}
}
