package version

import (
	"testing"
	"time"
)

func Test_default(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "provided", want: "one"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _default("one", "two"); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "default", want: "main.go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "default", want: "v0.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Version(); got != tt.want {
				t.Errorf("Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "default", want: time.Unix(0, 0).Format(time.RFC1123Z)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildDate(); got != tt.want {
				t.Errorf("BuildDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildType(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "default", want: "workstation"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildType(); got != tt.want {
				t.Errorf("BuildType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "default", want: "00000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildId(); got != tt.want {
				t.Errorf("BuildId() = %v, want %v", got, tt.want)
			}
		})
	}
}
