package grug

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func init() {
	AllActions = append(AllActions, []Action{
		{
			// Send a HTTP Get request
			// If a second arg is provided it will be used as the request body
			Name: "HTTPGet",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				url, ok := args[0].(string)
				if !ok {
					return nil, fmt.Errorf("unable to interpret %v as a string", args[0])
				}

				var resp *http.Response
				var err error
				if len(args) == 1 {
					resp, err = http.Get(url)
					if err != nil {
						return nil, err
					}
				} else {
					reqBody, ok := args[1].(string)
					if !ok {
						return nil, fmt.Errorf("unable to interpret %v as a string", args[1])
					}

					client := &http.Client{}
					req, err := http.NewRequest("GET", url, bytes.NewBufferString(reqBody))
					if err != nil {
						return nil, err
					}

					resp, err = client.Do(req)
					if err != nil {
						return nil, err
					}
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				}

				return body, nil
			},
		},
	}...)
}
