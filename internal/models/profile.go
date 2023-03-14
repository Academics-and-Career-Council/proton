package models

type candidate struct {
	Key  int `json:"key" bson:"key"`
	Course string `json:"course" bson:"course"`
	Grade   string `json:"grade" bson:"grade"`
	Credits  int `json:"credits" bson:"credits"`
	Credits_received float32 `json:"credits_received" bson:"credits_received"`
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`

    
}

type Book struct {
	ID     string `json:"id" bson:"_id"`
	Email  string `json:"email" bson:"email"`
	UserId string `json:"userid" bson:"userid"`
	Gradesdata  [][]candidate `json:"gradesData" bson:"gradesData"`
}

type Courses struct{
	Key  int `json:"key" bson:"key"`
	Course  string `json:"course" bson:"course"` 
	Value string `json:"value" bson:"value"`
	Cred  int `json:"cred" bson:"cred"`
	Category string `json:"category" bson:"category"`
	Grade  string `json:"grade" bson:"grade"`
	Credits_received  int `json:"credits_received" bson:"credits_received"` 
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`
	
	
}

type CoursesSem struct{
	Key  int `json:"key" bson:"key"`
	Course  string `json:"course" bson:"course"` 
	Value string `json:"value" bson:"value"`
	Cred  int `json:"credits" bson:"cred"`
	Category string `json:"category" bson:"category"`
	Grade  string `json:"grade" bson:"grade"`
	Sem  string `json:"sem" bson:"sem"`
	Branch string `json:"branch" bson:"branch"`
	Credits_received  int `json:"credits_received" bson:"credits_received"` 
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`
	
	
}

type CoursesY22 struct{
	Key  int `json:"key" bson:"key"`
	Course  string `json:"course" bson:"course"` 
	Credits int `json:"credits" bson:"credits"`
    Category string `json:"category" bson:"category"`
	Grade  string `json:"grade" bson:"grade"`
	Credits_received  int `json:"credits_received" bson:"credits_received"` 
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`
}

type CourseBranch struct{
	Branch string `json:"branch" bson:"branch"`
	Sem [][]courseBr `json:"sem" bson:"sem"`
}


type courseBr struct {
	Course string `json:"course" bson:"course"`
	Cred  string `json:"credits" bson:"cred"`
	Credits_received float32 `json:"credits_received" bson:"credits_received"`
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`

    
}