package service

import (
	"app/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func InsertEvent(c *gin.Context) error {
	title := c.PostForm("title")
	body := c.PostForm("body")
	eventDateStr := c.PostForm("eventDate")
	joinDeadlineDateStr := c.PostForm("joinDeadlineDate")
	capacityStr := c.PostForm("capacity")

	eventDate, err := time.Parse(time.RFC3339, eventDateStr)
	if err != nil {
		return err
	}

	joinDeadlineDate, err := time.Parse(time.RFC3339, joinDeadlineDateStr)
	if err != nil {
		return err
	}

	capacity, err := strconv.Atoi(capacityStr)
	if err != nil {
		return err
	}

	event := model.Event{
		Title:            title,
		Body:             body,
		EventDate:        eventDate,
		JoinDeadlineDate: joinDeadlineDate,
		Capacity:         capacity,
	}

	err = event.InsertEvent()
	if err != nil {
		return err
	}

	return nil
}