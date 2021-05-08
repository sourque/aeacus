package cmd

import (
	"time"
)

// metaConfig is the overarching context used by most functions in aeacus.
type metaConfig struct {
	DirPath     string
	TeamID      string
	Config      scoringChecks
	Image       imageData
	Conn        connData
	Connection  bool
	ShellActive bool
}

// imageData is the current scoring data for the image. It is able to be
// wiped, removed, etc, on each run without affecting anything else.
type imageData struct {
	RunningTime time.Time
	Contribs    int
	Detracts    int
	Score       int
	ScoredVulns int
	TotalPoints int
	Penalties   []scoreItem
	Points      []scoreItem
}

// connData represents the current connectivity state of the image to the
// internet and the scoring server.
type connData struct {
	OverallColor  string
	OverallStatus string
	NetColor      string
	NetStatus     string
	ServerColor   string
	ServerStatus  string
}

// scoreItem is the scoring report representation of a check, containing only
// the message and points associated with it.
type scoreItem struct {
	Message string
	Points  int
}

// scoringChecks is a representation of the TOML configuration typically
// specific in scoring.conf.
type scoringChecks struct {
	DisableShell bool
	Local        bool
	NoDestroy    bool
	EndDate      string
	Name         string
	OS           string
	Password     string
	Remote       string
	Title        string
	User         string
	Version      string
	Check        []check
}

// check is the smallest unit that can show up on a scoring report. It holds
// all the conditions for a check, and its message and points (autogenerated or
// otherwise).
type check struct {
	Message      string
	Points       int
	Fail         []condition
	Pass         []condition
	PassOverride []condition
}

// condition is a single pass/fail condition inside of a check. It supports up
// to four arguments.
type condition struct {
	Type string
	Arg1 string
	Arg2 string
	Arg3 string
	Arg4 string
}

// statusRes is to parse a JSON response from the remote server
type statusRes struct {
	Status string `json:"status"`
}