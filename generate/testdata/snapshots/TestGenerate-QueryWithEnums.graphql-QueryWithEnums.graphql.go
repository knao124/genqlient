// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// QueryWithEnumsOtherUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithEnumsOtherUser struct {
	Roles []Role `json:"roles"`
}

// GetRoles returns QueryWithEnumsOtherUser.Roles, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsOtherUser) GetRoles() []Role { return v.Roles }

// QueryWithEnumsResponse is returned by QueryWithEnums on success.
type QueryWithEnumsResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithEnumsUser `json:"user"`
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	OtherUser QueryWithEnumsOtherUser `json:"otherUser"`
}

// GetUser returns QueryWithEnumsResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsResponse) GetUser() QueryWithEnumsUser { return v.User }

// GetOtherUser returns QueryWithEnumsResponse.OtherUser, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsResponse) GetOtherUser() QueryWithEnumsOtherUser { return v.OtherUser }

// QueryWithEnumsUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithEnumsUser struct {
	Roles []Role `json:"roles"`
}

// GetRoles returns QueryWithEnumsUser.Roles, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsUser) GetRoles() []Role { return v.Roles }

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

func QueryWithEnums(
	client graphql.Client,
) (*QueryWithEnumsResponse, error) {
	req := &graphql.Request{
		OpName: "QueryWithEnums",
		Query: `
query QueryWithEnums {
	user {
		roles
	}
	otherUser: user {
		roles
	}
}
`,
	}
	var err error

	var data QueryWithEnumsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

