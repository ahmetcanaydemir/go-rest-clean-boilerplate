package model

import (
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/utils/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Stats struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Date  *helper.JsonTime   `json:"date"`
	Count float64            `json:"count"`
}
