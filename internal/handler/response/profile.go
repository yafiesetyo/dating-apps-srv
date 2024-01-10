package response

type Profile struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	Gender       string   `json:"gender"`
	ImageUrl     []string `json:"image_url"`
	DOB          string   `json:"dob"`
	POB          string   `json:"pob"`
	Religion     string   `json:"religion"`
	Description  string   `json:"description"`
	Hobby        string   `json:"hobby"`
	VerifiedUser bool     `json:"verified_user"`
}
