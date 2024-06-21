package version

import "time"

var (
	LName      string
	LVersion   string
	LBuildDate string
	LBuildType string
	LBuildId   string
)

func _default(val string, def string) string {
	if len(val) > 0 {
		return val
	}
	return def
}

func Name() string {
	return _default(LName, "main.go")
}

func Version() string {
	return _default(LVersion, "v0.0.0")
}

func BuildDate() string {
	return _default(LBuildDate, time.Unix(0, 0).Format(time.RFC1123Z))
}

func BuildType() string {
	return _default(LBuildType, "workstation")
}

func BuildId() string {
	return _default(LBuildId, "00000000")
}
