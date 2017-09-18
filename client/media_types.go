// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "celler": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/ReSTARTR/goa-sample/design
// --out=$(GOPATH)/src/github.com/ReSTARTR/goa-sample
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// user media type (default view)
//
// Identifier: application.vnd.user+json; view=default
type ApplicationVndUser struct {
	// description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// DecodeApplicationVndUser decodes the ApplicationVndUser instance encoded in resp body.
func (c *Client) DecodeApplicationVndUser(resp *http.Response) (*ApplicationVndUser, error) {
	var decoded ApplicationVndUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A bottle of wine (default view)
//
// Identifier: application/vnd.goa.example.com+json; view=default
type GoaExampleCom struct {
	// API href for making requests on the bottle
	Href string `form:"href" json:"href" xml:"href"`
	// Unique bottle ID
	ID int `form:"id" json:"id" xml:"id"`
	// Name of wine
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GoaExampleCom media type instance.
func (mt *GoaExampleCom) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeGoaExampleCom decodes the GoaExampleCom instance encoded in resp body.
func (c *Client) DecodeGoaExampleCom(resp *http.Response) (*GoaExampleCom, error) {
	var decoded GoaExampleCom
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
