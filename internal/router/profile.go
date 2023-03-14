package router

import (
	"proton/internal/common"
	"proton/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "strconv"
)

func AddcourseGroup(app *fiber.App) {
	courseGroup := app.Group("/")

	// courseGroup.Get("/", getBooks)
	courseGroup.Get("/courses", getCourses)
	courseGroup.Get("/courses1/:category", getCoursesY22)
	// courseGroup.Get("/getcourse/:key", getCourse)
	courseGroup.Get("/getuser/:email", getProfile)
	courseGroup.Post("/register1", createProfile)
	courseGroup.Post("/register2", createCourse)
	courseGroup.Put("/updateuser/:email", updateProfile)
	courseGroup.Put("/updatecourse/:course", updateCourse)
	courseGroup.Get("/coursessem/:branch/:sem", getCoursesSem)
	courseGroup.Get("/coursebranch/:branch", getCoursesBranch)

	

	// courseGroup.Delete("/:id", deleteBook)
}


type coursessemDTO struct{
	Key  int `json:"key" bson:"key"`
	Course string `json:"course" bson:"course"` 
	Calue string `json:"value" bson:"value"`
	Cred int `json:"cred" bson:"cred"`
	Category string `json:"category" bson:"category"`
	Grade  string `json:"grade" bson:"grade"`
	Sem  string `json:"sem" bson:"sem"`
	Branch string `json:"branch" bson:"branch"`
	Credits_received  int `json:"credits_received" bson:"credits_received"` 
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`
}

type coursesDTO struct{
	Key  int `json:"key" bson:"key"`
	Course string `json:"course" bson:"course"` 
	Calue string `json:"value" bson:"value"`
	Cred int `json:"credits" bson:"cred"`
	Category string `json:"category" bson:"category"`
	Grade  string `json:"grade" bson:"grade"`
	Sem  string `json:"sem" bson:"sem"`
	Branch string `json:"branch" bson:"branch"`
	Credits_received  int `json:"credits_received" bson:"credits_received"` 
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`
}

func getCoursesSem(c *fiber.Ctx) error {
	coll := common.GetDBCollection("courses")

	sem := c.Params("sem")
	branch:=c.Params("branch")
	if (branch=="") {
		return c.Status(400).JSON(fiber.Map{
			"error": "sem and branch is required",
			"sem": c.Params("sem"),
			"branch":c.Params("branch"),
		})
	}
	// find all books
	books := make([]models.CoursesSem, 0)
	cursor, err := coll.Find(c.Context(), bson.M{"branch":branch,"sem":sem})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		book := models.CoursesSem{}
		err := cursor.Decode(&book)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		books = append(books, book)
	}

	return c.Status(200).JSON(fiber.Map{"data": books,"sem":sem,"branch":branch})
}


func getCourses(c *fiber.Ctx) error {
	coll := common.GetDBCollection("courses")

	// find all books
	books := make([]models.Courses, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		book := models.Courses{}
		err := cursor.Decode(&book)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		books = append(books, book)
	}

	return c.Status(200).JSON(fiber.Map{"data": books})
}
type coursesY22DTO struct{
	Key  int `json:"key" bson:"key"`
	Course  string `json:"course" bson:"course"` 
	Credits  int `json:"credits" bson:"credits"`
    Category string `json:"category" bson:"category"`
	Grade  string `json:"grade" bson:"grade"`
	Credits_received  int `json:"credits_received" bson:"credits_received"` 
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`
}
func getCoursesY22(c *fiber.Ctx) error {
	coll := common.GetDBCollection("courses")

	category := c.Params("category")
	if category == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "category is required",
			"category": c.Params("category"),
		})
	}
	// find all books
	
	books := make([]models.CoursesY22, 0)
	cursor, err := coll.Find(c.Context(), bson.M{"Category":category})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		book := models.CoursesY22{}
		err := cursor.Decode(&book)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		books = append(books, book)
	}

	return c.Status(200).JSON(fiber.Map{"data": books})
}

func getProfile(c *fiber.Ctx) error {
	coll := common.GetDBCollection("books")

	
	email := c.Params("email")
	if email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "email is required",
		})
	}
	

	book := models.Book{}

	err := coll.FindOne(c.Context(), bson.M{"email": email}).Decode(&book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": book})
}


type candidate struct {
	Key  int `json:"key" bson:"key"`
	Course string `json:"course" bson:"course"`
	Grade   string `json:"grade" bson:"grade"`
	Credits  int `json:"credits" bson:"credits"`
	Credits_received float32 `json:"credits_received" bson:"credits_received"`
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`   
}

type createDTO struct {
	Email  string `json:"email" bson:"email"`
	UserId string `json:"userid" bson:"userid"`
	Gradesdata  [][]candidate `json:"gradesData" bson:"gradesData"`
}


func createProfile(c *fiber.Ctx) error {
	// validate the body
	b := new(createDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// create the book
	coll := common.GetDBCollection("books")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create book",
			"message": err.Error(),
		})
	}

	// return the book
	return c.Status(201).JSON(fiber.Map{
		"result": result,
		"result1":c.Context(),
	})
}
type courseBr struct {
	Course string `json:"course" bson:"course"`
	Cred  string `json:"credits" bson:"cred"`
	Credits_received float32 `json:"credits_received" bson:"credits_received"`
	Is_repeated   bool `json:"is_repeated" bson:"is_repeated"`
	Is_sx  bool `json:"is_sx" bson:"is_sx"`

    
}
type courseBranchDTO struct{
	Branch string `json:"branch" bson:"branch"`
	Sem [][]courseBr `json:"sem" bson:"sem"`
}

func getCoursesBranch(c *fiber.Ctx) error {
	// validate the body
	coll := common.GetDBCollection("coursebranch")

	branch := c.Params("branch")
	// sem := c.Params("sem")
	if branch == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "branch is required",
		})
	}
	

	book := models.CourseBranch{}

	err := coll.FindOne(c.Context(), bson.M{"branch": branch}).Decode(&book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": book})
}
// type coursesDTO struct{
// 	Key  int `json:"key" bson:"key"`
// 	Course string `json:"course" bson:"course"` 
// 	Calue string `json:"value" bson:"value"`
// 	Cred int `json:"cred" bson:"cred"`
// }

func createCourse(c *fiber.Ctx) error {
	// validate the body
	b := new(coursesDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// create the book
	coll := common.GetDBCollection("courses")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create book",
			"message": err.Error(),
		})
	}

	// return the book
	return c.Status(201).JSON(fiber.Map{
		"result1": c.Context(),
		"result":result,
		"b":b,
	})
}


// type updateDTO struct {
// 	Title  string `json:"title,omitempty" bson:"title,omitempty"`
// 	Author string `json:"author,omitempty" bson:"author,omitempty"`
// 	Year   string `json:"year,omitempty" bson:"year,omitempty"`
// 	// Gradesdata  [][3]string `json:"gradesdata" bson:"gradesdata"`
// 	Gradesdatas  [][]candidate `json:"gradesdatas" bson:"gradesdatas"`
// }
type updateDTO struct {
	Email  string `json:"email" bson:"email"`
	UserId string `json:"userid" bson:"userid"`
	Gradesdata  [][]candidate `json:"gradesData" bson:"gradesData"`
}


func updateCourse(c *fiber.Ctx) error {
	// validate the body
	b := new(coursesDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// get the id
	course := c.Params("course")
	if course == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "email is required",
		})
	}
	// objectId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error": "invalid id",
	// 	})
	// }

	// update the book
	coll := common.GetDBCollection("courses")
	result, err := coll.UpdateOne(c.Context(), bson.M{"course": course}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update book",
			"course":course,
			"message": err.Error(),
		})
	}

	// return the book
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}


func updateProfile(c *fiber.Ctx) error {
	// validate the body
	b := new(updateDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// get the id
	email := c.Params("email")
	if email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "email is required",
		})
	}
	// objectId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error": "invalid id",
	// 	})
	// }

	// update the book
	coll := common.GetDBCollection("books")
	result, err := coll.UpdateOne(c.Context(), bson.M{"email": email}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update book",
			"email":email,
			"message": err.Error(),
		})
	}

	// return the book
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

// func deleteBook(c *fiber.Ctx) error {
// 	// get the id
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "invalid id",
// 		})
// 	}

// 	// delete the book
// 	coll := common.GetDBCollection("books")
// 	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error":   "Failed to delete book",
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"result": result,
// 	})
// }
