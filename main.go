package main
func translateTime(timeStr string, toTZ string) (string, error) {
	// Kindly provided and debugged by and with Claude
	// Don't worry, it won't explode, I already tested it ;)

	// Load both Local and Tokyo timezones
	fromLoc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}

	toLoc, err := time.LoadLocation(toTZ)
	if err != nil {
		return "", err
	}

	// FIX: Incorrect time by 01:04:00
	// Why the fuck are they "historically accurate"? They think I'm gonna travel
	// back in time???

	now := time.Now().In(fromLoc)
	datePrefix := now.Format("2006-01-02")

	// Convert our string into LocalTimezone
	t, err := time.ParseInLocation("2006-01-02 15:04", datePrefix+" "+timeStr, fromLoc)
	if err != nil {
		return "", err
	}

	// Finally convert our LocalTime obj to Tokyo
	return t.In(toLoc).Format("15:04"), nil
}

func at(run_time string) {
	// Written with Gemini and Claude because "at" is a special snowflake
	// that won't accept arguments beyond time, won't drop you into stdin
	// unless you're in your terminal (as in your terminal prompt),
	// so you can really only pass the command this way from here

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	input, err := line.Prompt("> ")

	if err != nil {
		if err == liner.ErrPromptAborted {
			log.Fatal("Aborted")
		}
		log.Printf("Error reading line: %v", err)
	}

	cmd := exec.Command("at", run_time)
	cmd.Stdin = strings.NewReader(input)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	log.Print(string(output))
}

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

		log.Printf("Hours: %s Minutes: %s\n", hours, minutes)

		timeToParse := fmt.Sprintf("%s:%s", hours, minutes)
		timeParsed, err := translateTime(timeToParse, "America/Mexico_City")

		if err != nil {
			log.Fatal(err)
		}

		log.Println("JST:", timeToParse, "Local:", timeParsed)

		at(timeParsed)

	case re_special.MatchString(run_hour):
		matches := re_special.FindStringSubmatch(run_hour)
		special := matches[1]

		var hours string

		switch special {
		case "noon":
			hours = "12"
		case "teatime":
			hours = "16"
		case "midnight":
			hours = "24"
		}

		minutes := "00"

		log.Printf("Special: %s (Hours: %s Minutes: %s)\n", special, hours, minutes)

		timeToParse := fmt.Sprintf("%s:%s", hours, minutes)
		timeParsed, err := translateTime(timeToParse, "America/Mexico_City")

		if err != nil {
			log.Fatal(err)
		}

		log.Println("JST:", timeToParse, "Local:", timeParsed)

	}
}
