// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "user": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-user/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-user
// --version=v1.3.1

package app

import (
	"github.com/keitaroinc/goa"
)

// users media type (default view)
//
// Identifier: application/vnd.goa.user+json; view=default
type Users struct {
	// Status of user account
	Active bool `form:"active" json:"active" yaml:"active" xml:"active"`
	// Email of user
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// External id of user
	ExternalID string `form:"externalId" json:"externalId" yaml:"externalId" xml:"externalId"`
	// Unique user ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// List of namespaces this user belongs to
	Namespaces []string `form:"namespaces,omitempty" json:"namespaces,omitempty" yaml:"namespaces,omitempty" xml:"namespaces,omitempty"`
	// List of organizations to which this user belongs to
	Organizations []string `form:"organizations,omitempty" json:"organizations,omitempty" yaml:"organizations,omitempty" xml:"organizations,omitempty"`
	// Roles of user
	Roles []string `form:"roles" json:"roles" yaml:"roles" xml:"roles"`
}

// Validate validates the Users media type instance.
func (mt *Users) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Roles == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "roles"))
	}
	if mt.ExternalID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "externalId"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	return
}

// ResetToken media type (default view)
//
// Identifier: resettokenmedia; view=default
type ResetToken struct {
	// User email
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// User ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// New token
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
}

// Validate validates the ResetToken media type instance.
func (mt *ResetToken) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}
