//go:build ignore
// +build ignore

package main

import (
	"github.com/lunelabs/go-gen-rest/gen"
	"github.com/lunelabs/go-gen-rest/model"
)

func main() {
	gen.NewRest(
		"github.com/lunelabs/go-gen-rest",
		[]model.Resource{
			model.Resource{
				Name: "ip_list",
				Fields: []model.Field{
					{
						Name:    "hash",
						Type:    "string",
						Filter:  false,
						IdField: true,
					},
					{
						Name:            "title",
						Type:            "string",
						Filter:          true,
						IdField:         false,
						CreateValidator: "required",
						FilterValidator: "required",
					},
				},
			},
		},
	).Generate()
}
