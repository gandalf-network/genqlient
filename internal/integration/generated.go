package integration

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

// queryWithInterfaceNoFragmentsBeing includes the requested fields of the GraphQL type Being.
type queryWithInterfaceNoFragmentsBeing interface {
	implementsGraphQLInterfacequeryWithInterfaceNoFragmentsBeing()
}

func (v *queryWithInterfaceNoFragmentsBeingUser) implementsGraphQLInterfacequeryWithInterfaceNoFragmentsBeing() {
}
func (v *queryWithInterfaceNoFragmentsBeingAnimal) implementsGraphQLInterfacequeryWithInterfaceNoFragmentsBeing() {
}

func __unmarshalqueryWithInterfaceNoFragmentsBeing(v *queryWithInterfaceNoFragmentsBeing, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "User":
		*v = new(queryWithInterfaceNoFragmentsBeingUser)
		return json.Unmarshal(m, *v)
	case "Animal":
		*v = new(queryWithInterfaceNoFragmentsBeingAnimal)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for queryWithInterfaceNoFragmentsBeing: "%v"`, tn.TypeName)
	}
}

// queryWithInterfaceNoFragmentsBeingAnimal includes the requested fields of the GraphQL type Animal.
type queryWithInterfaceNoFragmentsBeingAnimal struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceNoFragmentsBeingUser includes the requested fields of the GraphQL type User.
type queryWithInterfaceNoFragmentsBeingUser struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceNoFragmentsMeUser includes the requested fields of the GraphQL type User.
type queryWithInterfaceNoFragmentsMeUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// queryWithInterfaceNoFragmentsResponse is returned by queryWithInterfaceNoFragments on success.
type queryWithInterfaceNoFragmentsResponse struct {
	Being queryWithInterfaceNoFragmentsBeing  `json:"-"`
	Me    queryWithInterfaceNoFragmentsMeUser `json:"me"`
}

func (v *queryWithInterfaceNoFragmentsResponse) UnmarshalJSON(b []byte) error {

	type queryWithInterfaceNoFragmentsResponseWrapper queryWithInterfaceNoFragmentsResponse

	var firstPass struct {
		*queryWithInterfaceNoFragmentsResponseWrapper
		Being json.RawMessage `json:"being"`
	}
	firstPass.queryWithInterfaceNoFragmentsResponseWrapper = (*queryWithInterfaceNoFragmentsResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = __unmarshalqueryWithInterfaceNoFragmentsBeing(
		&v.Being, firstPass.Being)
	if err != nil {
		return err
	}

	return nil
}

// queryWithVariablesResponse is returned by queryWithVariables on success.
type queryWithVariablesResponse struct {
	User queryWithVariablesUser `json:"user"`
}

// queryWithVariablesUser includes the requested fields of the GraphQL type User.
type queryWithVariablesUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LuckyNumber int    `json:"luckyNumber"`
}

// simpleQueryMeUser includes the requested fields of the GraphQL type User.
type simpleQueryMeUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LuckyNumber int    `json:"luckyNumber"`
}

// simpleQueryResponse is returned by simpleQuery on success.
type simpleQueryResponse struct {
	Me simpleQueryMeUser `json:"me"`
}

func simpleQuery(
	ctx context.Context,
	client graphql.Client,
) (*simpleQueryResponse, error) {
	var retval simpleQueryResponse
	err := client.MakeRequest(
		ctx,
		"simpleQuery",
		`
query simpleQuery {
	me {
		id
		name
		luckyNumber
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

func queryWithVariables(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*queryWithVariablesResponse, error) {
	variables := map[string]interface{}{
		"id": id,
	}

	var retval queryWithVariablesResponse
	err := client.MakeRequest(
		ctx,
		"queryWithVariables",
		`
query queryWithVariables ($id: ID!) {
	user(id: $id) {
		id
		name
		luckyNumber
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}

func queryWithInterfaceNoFragments(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*queryWithInterfaceNoFragmentsResponse, error) {
	variables := map[string]interface{}{
		"id": id,
	}

	var retval queryWithInterfaceNoFragmentsResponse
	err := client.MakeRequest(
		ctx,
		"queryWithInterfaceNoFragments",
		`
query queryWithInterfaceNoFragments ($id: ID!) {
	being(id: $id) {
		__typename
		id
		name
	}
	me {
		id
		name
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}
