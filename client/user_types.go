// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "user": Application User Types
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/user-microservice/design
// --out=$(GOPATH)/src/github.com/JormungandrK/user-microservice
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// Email and password credentials
type credentials struct {
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the credentials type instance.
func (ut *credentials) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 6, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 30, false))
		}
	}
	return
}

// Publicize creates Credentials from credentials
func (ut *credentials) Publicize() *Credentials {
	var pub Credentials
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// Email and password credentials
type Credentials struct {
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
	// Password of user
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate validates the Credentials type instance.
func (ut *Credentials) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if utf8.RuneCountInString(ut.Password) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 6, true))
	}
	if utf8.RuneCountInString(ut.Password) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 30, false))
	}
	return
}

// Email payload
type emailPayload struct {
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// Validate validates the emailPayload type instance.
func (ut *emailPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates EmailPayload from emailPayload
func (ut *emailPayload) Publicize() *EmailPayload {
	var pub EmailPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	return &pub
}

// Email payload
type EmailPayload struct {
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// Validate validates the EmailPayload type instance.
func (ut *EmailPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	return
}

// UserPayload
type userPayload struct {
	// Status of user account
	Active *bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// External id of user
	ExternalID *string `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Roles of user
	Roles []string `form:"roles,omitempty" json:"roles,omitempty" xml:"roles,omitempty"`
	// Token for email verification
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// Finalize sets the default values for userPayload type instance.
func (ut *userPayload) Finalize() {
	var defaultActive = false
	if ut.Active == nil {
		ut.Active = &defaultActive
	}
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 6, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 30, false))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Active != nil {
		pub.Active = *ut.Active
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.ExternalID != nil {
		pub.ExternalID = ut.ExternalID
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.Roles != nil {
		pub.Roles = ut.Roles
	}
	if ut.Token != nil {
		pub.Token = ut.Token
	}
	return &pub
}

// UserPayload
type UserPayload struct {
	// Status of user account
	Active bool `form:"active" json:"active" xml:"active"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
	// External id of user
	ExternalID *string `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Roles of user
	Roles []string `form:"roles,omitempty" json:"roles,omitempty" xml:"roles,omitempty"`
	// Token for email verification
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 6, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 30, false))
		}
	}
	return
}
