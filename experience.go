package draven

type Experience struct {
	ID          string
	Description string
	From        User
	Name        string
	With        []User
}

type EducationExperience struct {
	ID            string
	Classes       []Experience
	Concentration []Page
	Degree        Page
	School        Page
	Type          string
	With          []User
	Year          Page
}
