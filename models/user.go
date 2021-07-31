package models

type User struct {
	Id               uint   `json:"id" valid:"-"`
	FirstName        string `json:"first_name" valid:"required~First name is required"`
	LastName         string `json:"last_name" valid:"required~Last name is required"`
	Email            string `json:"email" gorm:"unique" valid:"email"`
	Password         []byte `json:"-" valid:"required~Password is required"`
	Phone            string `json:"phone" valid:"required~Phone is required"`
	Image            string `json:"image" valid:"-"`
	AlternativePhone string `json:"alternative_phone" valid:"-"`
	Country          string `json:"country" valid:"required~Country is required"`
	Idnumber         string `json:"id_number" valid:"required~Id number is required"`
	Dob              string `json:"dob" valid:"required~Date of birth is required"`
	Gender           string `json:"gender" valid:"required~Gender is required"`
	MaritalStatus    string `json:"-" valid:"required~Marital status is required"`
	Address          string `json:"-" valid:"required~Address is required"`
	Latitude         string `json:"-" valid:"required~Latitude is required"`
	Longitude        string `json:"-" valid:"required~Longitude is required"`
	JobType          string `json:"job_type" valid:"required~Job type is required"`
	Referee          string `json:"referee" valid:"-"`
	Comapny          string `json:"-" valid:"required~Company is required"`
	Position         string `json:"-" valid:"required~Position is required"`
	CompanyPhone     string `json:"-" valid:"required~Company phone is required"`
	IncomeRange      string `json:"-" valid:"required~Income range is required"`
	LoginPin         []byte `json:"-" gorm:"default:null" valid:"-"`
}
