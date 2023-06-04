package gofunctions

import (
	"fmt"
	"net/http"
)

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// pointer receivers
func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
}

func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1
}

//unsemantic

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}
