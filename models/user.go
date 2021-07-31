package models

type User struct {
	Id               uint   `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email" gorm:"unique"`
	Password         []byte `json:"-"`
	Phone            string `json:"phone"`
	Image            string `json:"image" gorm:"default:null"`
	AlternativePhone string `json:"alternative_phone"`
	Country          string `json:"country"`
	Idnumber         string `json:"id_number"`
	Dob              string `json:"dob"`
	Gender           string `json:"gender"`
	MaritalStatus    string `json:"-"`
	Address          string `json:"-"`
	Latitude         string `json:"-"`
	Longitude        string `json:"-"`
	JobType          string `json:"job_type"`
	Referee          string `json:"referee"`
	Comapny          string `json:"-"`
	Position         string `json:"-"`
	CompanyPhone     string `json:"-"`
	IncomeRange      string `json:"-"`
	LoginPin         []byte `json:"-" gorm:"default:null"`
}
