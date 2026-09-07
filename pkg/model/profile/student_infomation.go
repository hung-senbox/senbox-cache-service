package profile

type StudentInformation struct {
	ID                string `json:"id"`
	OrganizationId    string `json:"organization_id"`
	StudentId         string `json:"student_id"`
	DOB               string `json:"dob" bson:"dob"`
	Gender            uint   `json:"gender" bson:"gender"`
	StudyLevel        uint   `json:"study_level" bson:"study_level"`
	MinWaterMustDrink uint   `json:"min_water_must_drink" bson:"min_water_must_drink"`
	Description       string `json:"description" bson:"description"`
	Mode              string `json:"mode" bson:"mode"`
	CommunicatorLevel uint   `json:"communicator_level" bson:"communicator_level"`
}
