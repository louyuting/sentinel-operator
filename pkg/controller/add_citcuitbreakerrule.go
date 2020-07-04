package controller

import (
	"github.com/alibaba/sentinel-golang/pkg/controller/citcuitbreakerrule"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, citcuitbreakerrule.Add)
}
