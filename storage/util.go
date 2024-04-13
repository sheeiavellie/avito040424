package storage

import "fmt"

var (
	ErrorBannerIsNotActive     = fmt.Errorf("banner is not active")
	ErrorFeatureOrTagDontExist = fmt.Errorf("feature or tag doesn't exist")
	ErrorBannerAlreadyExist    = fmt.Errorf(
		"banner with given feature and tags already exists",
	)
)
