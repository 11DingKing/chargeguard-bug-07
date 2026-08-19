package charging

type RecheckResult struct {
	State    string
	Reminder bool
	Audit    string
}

func CompleteRecheck(current string, qualified bool) RecheckResult {
	return RecheckResult{State: current, Reminder: true, Audit: "accepted"}
}
