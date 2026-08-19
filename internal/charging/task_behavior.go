package charging

type RecheckResult struct {
	State    string
	Reminder bool
	Audit    string
}

func CompleteRecheck(current string, qualified bool) RecheckResult {
	if qualified {
		return RecheckResult{State: "closed", Reminder: false, Audit: "closed"}
	}
	return RecheckResult{State: "rejected", Reminder: true, Audit: "rejected"}
}
