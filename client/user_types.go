// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "user": Application User Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-user/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-user
// --version=v1.3.1

package client

import (
	"github.com/keitaroinc/goa"
	"unicode/utf8"
)

// CreateUserPayload
type createUserPayload struct {
	// Status of user account
	Active *bool `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// External id of user
	ExternalID *string `form:"externalId,omitempty" json:"externalId,omitempty" yaml:"externalId,omitempty" xml:"externalId,omitempty"`
	// List of namespaces this user belongs to
	Namespaces []string `form:"namespaces,omitempty" json:"namespaces,omitempty" yaml:"namespaces,omitempty" xml:"namespaces,omitempty"`
	// List of organizations to which this user belongs to
	Organizations []string `form:"organizations,omitempty" json:"organizations,omitempty" yaml:"organizations,omitempty" xml:"organizations,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	// Roles of user
	Roles []string `form:"roles,omitempty" json:"roles,omitempty" yaml:"roles,omitempty" xml:"roles,omitempty"`
	// Token for email verification
	Token *string `form:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty" xml:"token,omitempty"`
}

// Finalize sets the default values for createUserPayload type instance.
func (ut *createUserPayload) Finalize() {
	var defaultActive = false
	if ut.Active == nil {
		ut.Active = &defaultActive
	}
}

// Validate validates the createUserPayload type instance.
func (ut *createUserPayload) Validate() (err error) {
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

// Publicize creates CreateUserPayload from createUserPayload
func (ut *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if ut.Active != nil {
		pub.Active = *ut.Active
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.ExternalID != nil {
		pub.ExternalID = ut.ExternalID
	}
	if ut.Namespaces != nil {
		pub.Namespaces = ut.Namespaces
	}
	if ut.Organizations != nil {
		pub.Organizations = ut.Organizations
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

// CreateUserPayload
type CreateUserPayload struct {
	// Status of user account
	Active bool `form:"active" json:"active" yaml:"active" xml:"active"`
	// Email of user
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// External id of user
	ExternalID *string `form:"externalId,omitempty" json:"externalId,omitempty" yaml:"externalId,omitempty" xml:"externalId,omitempty"`
	// List of namespaces this user belongs to
	Namespaces []string `form:"namespaces,omitempty" json:"namespaces,omitempty" yaml:"namespaces,omitempty" xml:"namespaces,omitempty"`
	// List of organizations to which this user belongs to
	Organizations []string `form:"organizations,omitempty" json:"organizations,omitempty" yaml:"organizations,omitempty" xml:"organizations,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	// Roles of user
	Roles []string `form:"roles,omitempty" json:"roles,omitempty" yaml:"roles,omitempty" xml:"roles,omitempty"`
	// Token for email verification
	Token *string `form:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty" xml:"token,omitempty"`
}

// Validate validates the CreateUserPayload type instance.
func (ut *CreateUserPayload) Validate() (err error) {
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

// Email and password credentials
type credentials struct {
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
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
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// Password of user
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
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
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
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
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
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

// Password Reset payload
type forgotPasswordPayload struct {
	// Email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// New password
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	// Forgot password token
	Token *string `form:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty" xml:"token,omitempty"`
}

// Validate validates the forgotPasswordPayload type instance.
func (ut *forgotPasswordPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Token == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "token"))
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

// Publicize creates ForgotPasswordPayload from forgotPasswordPayload
func (ut *forgotPasswordPayload) Publicize() *ForgotPasswordPayload {
	var pub ForgotPasswordPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.Token != nil {
		pub.Token = *ut.Token
	}
	return &pub
}

// Password Reset payload
type ForgotPasswordPayload struct {
	// Email of the user
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// New password
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
	// Forgot password token
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
}

// Validate validates the ForgotPasswordPayload type instance.
func (ut *ForgotPasswordPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if ut.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "token"))
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

// UpdateUserPayload
type updateUserPayload struct {
	// Status of user account
	Active *bool `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// External id of user
	ExternalID *string `form:"externalId,omitempty" json:"externalId,omitempty" yaml:"externalId,omitempty" xml:"externalId,omitempty"`
	// List of namespaces this user belongs to
	Namespaces []string `form:"namespaces,omitempty" json:"namespaces,omitempty" yaml:"namespaces,omitempty" xml:"namespaces,omitempty"`
	// List of organizations to which this user belongs to
	Organizations []string `form:"organizations,omitempty" json:"organizations,omitempty" yaml:"organizations,omitempty" xml:"organizations,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	// Roles of user
	Roles []string `form:"roles,omitempty" json:"roles,omitempty" yaml:"roles,omitempty" xml:"roles,omitempty"`
	// Token for email verification
	Token *string `form:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty" xml:"token,omitempty"`
}

// Finalize sets the default values for updateUserPayload type instance.
func (ut *updateUserPayload) Finalize() {
	var defaultActive = false
	if ut.Active == nil {
		ut.Active = &defaultActive
	}
}

// Validate validates the updateUserPayload type instance.
func (ut *updateUserPayload) Validate() (err error) {
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

// Publicize creates UpdateUserPayload from updateUserPayload
func (ut *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if ut.Active != nil {
		pub.Active = *ut.Active
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.ExternalID != nil {
		pub.ExternalID = ut.ExternalID
	}
	if ut.Namespaces != nil {
		pub.Namespaces = ut.Namespaces
	}
	if ut.Organizations != nil {
		pub.Organizations = ut.Organizations
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

// UpdateUserPayload
type UpdateUserPayload struct {
	// Status of user account
	Active bool `form:"active" json:"active" yaml:"active" xml:"active"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// External id of user
	ExternalID *string `form:"externalId,omitempty" json:"externalId,omitempty" yaml:"externalId,omitempty" xml:"externalId,omitempty"`
	// List of namespaces this user belongs to
	Namespaces []string `form:"namespaces,omitempty" json:"namespaces,omitempty" yaml:"namespaces,omitempty" xml:"namespaces,omitempty"`
	// List of organizations to which this user belongs to
	Organizations []string `form:"organizations,omitempty" json:"organizations,omitempty" yaml:"organizations,omitempty" xml:"organizations,omitempty"`
	// Password of user
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	// Roles of user
	Roles []string `form:"roles,omitempty" json:"roles,omitempty" yaml:"roles,omitempty" xml:"roles,omitempty"`
	// Token for email verification
	Token *string `form:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty" xml:"token,omitempty"`
}

// Validate validates the UpdateUserPayload type instance.
func (ut *UpdateUserPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, *ut.Email, goa.FormatEmail, err2))
		}
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
