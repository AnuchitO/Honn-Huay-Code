package skill

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Level struct {
	Key          string   `json:"key"`
	Name         string   `json:"name"`
	Brief        string   `json:"brief"`
	Descriptions []string `json:"descriptions"`
	Level        int      `json:"level"`
}

type Skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Levels      []Level  `json:"levels"`
	Tags        []string `json:"tags"`
}

type handler struct {
	storage storage
}

func NewHandler(db *sql.DB) handler {
	return handler{storage: storage{db: db}}
}

func (h handler) GetSkillByKey(c *gin.Context) {
	// get the key from the URL path param
	key := c.Param("key")

	skill, err := h.storage.findSkillByKey(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": skill})
}
