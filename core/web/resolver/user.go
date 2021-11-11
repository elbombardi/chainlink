package resolver

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/smartcontractkit/chainlink/core/sessions"
)

// UserResolver resolves the User type
type UserResolver struct {
	user *sessions.User
}

func NewUser(user *sessions.User) *UserResolver {
	return &UserResolver{user}
}

// Email resolves the user's email
func (r *UserResolver) Email() string {
	return r.user.Email
}

// CreatedAt resolves the user's creation date
func (r *UserResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.user.CreatedAt}
}

// -- UpdatePassword Mutation --

type UpdatePasswordInput struct {
	OldPassword string
	NewPassword string
}

// UpdatePasswordPayloadResolver resolves the payload type
type UpdatePasswordPayloadResolver struct {
	user *sessions.User
	// inputErrors maps an input path to a string
	inputErrs map[string]string
}

func NewUpdatePasswordPayload(user *sessions.User, inputErrs map[string]string) *UpdatePasswordPayloadResolver {
	return &UpdatePasswordPayloadResolver{user, inputErrs}
}

func (r *UpdatePasswordPayloadResolver) ToUpdatePasswordSuccess() (*UpdatePasswordSuccessResolver, bool) {
	if r.user == nil {
		return nil, false
	}

	return NewUpdatePasswordSuccess(r.user), true
}

func (r *UpdatePasswordPayloadResolver) ToInputErrors() (*InputErrorsResolver, bool) {
	if r.inputErrs != nil {
		var errs []*InputErrorResolver

		for path, message := range r.inputErrs {
			errs = append(errs, NewInputError(path, message))
		}

		return NewInputErrors(errs), true
	}

	return nil, false
}

type UpdatePasswordSuccessResolver struct {
	user *sessions.User
}

func NewUpdatePasswordSuccess(user *sessions.User) *UpdatePasswordSuccessResolver {
	return &UpdatePasswordSuccessResolver{user}
}

func (r *UpdatePasswordSuccessResolver) User() *UserResolver {
	return NewUser(r.user)
}
