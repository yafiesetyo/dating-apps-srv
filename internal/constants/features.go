package constants

type Features int

const (
	UnlimitedSwipe Features = iota + 1
	VerifiedUser
)

func GetFeature(in string) Features {
	switch in {
	case "unlimited_swipe":
		return UnlimitedSwipe
	case "verified_user":
		return VerifiedUser
	}

	return 0
}
