package apis

import (
	"encoding/json"
	"fmt"
	"gin/ginDemo/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context)  {
	c.String(http.StatusOK, "It works!")
}

func AddPerson(c *gin.Context)  {

	var user map[string]interface{}
	result, err := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(result, &user)
	firstName := user["first_name"].(string)
	lastName := user["last_name"].(string)

	p := models.Person{FirstName: firstName, LastName: lastName}
	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func ModPerson(c *gin.Context)  {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	p := models.Person{Id:id}

	p.GetPerson()

	if p.FirstName != "" {
		p.FirstName = firstName
		p.LastName = lastName
		ra, err := p.ModPerson()

		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprint("update successful %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

func DelPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	p := models.Person{Id: id}

	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func GetPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	person := models.Person{Id: id}

	person.GetPerson()
	if person.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": person,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

func GetPersonsApi(c *gin.Context)  {
	var p models.Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": persons,
		"msg":  "success",
	})
}
