package lead

import(
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ramgopalsiddh/crm_with_go-fiber/database"
	_"github.com/jinzhu/gorm/dialects/postgres"

)

type Lead struct{
	gorm.Model
	Name 	string `json:"name"`
	Company string `json:"company"`
	Email 	string `json:"email"`
	Phone 	string `json:"phone"`
}


func GetLeads(c *fiber.Ctx){
	db := database.DBconn
	var leads []Lead
	if err := db.Find(&leads).Error; err != nil {
		c.Status(500).Send(err)
		return
	}
	if err := c.JSON(leads); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBconn
	var lead Lead
	if err := db.Find(&lead, id).Error; err != nil {
		c.Status(500).Send(err)
		return
	}
	if err := c.JSON(lead); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func NewLead(c *fiber.Ctx){
	db := database.DBconn
	lead := new(Lead)
	if err := c.BodyParser(lead); err !=nil {
		c.Status(503).Send(err)
		return
	}
	if err := db.Create(&lead).Error; err != nil {
		c.Status(500).Send(err)
		return
	}
	if err := c.JSON(lead); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBconn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == ""{
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully Deleted")

}

