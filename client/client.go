package client

import (
	"encoding/json"
	"fmt"
	"go-inkdrop/config"
	"io/ioutil"
	"net/http"
	"net/url"
    "go-inkdrop/inkdrop"
)

type Parameter map[string]string


func Get[T any](param Parameter) (T, error) {
    var none T
    path := toPath(none)

    if path == "" {
        return none, fmt.Errorf("failed convert type to path : %T", none)
    }

	conf, err := config.Instance()
	if err != nil {
		return none, err
	}

	root := "http://" + conf.BindAddress + ":" + conf.Port
    // https://docs.inkdrop.app/manual/accessing-the-local-database#get-docid
    var endpoint string
    if path == "note" {
        docid, ok :=  param["docid"]
        if !ok {
            return none, fmt.Errorf("wrong parameter. Note needs docid in inkdrop.Parameter")
        }
        endpoint, err = url.JoinPath(root, docid)
    } else {
        endpoint, err = url.JoinPath(root, path)
    }

    if err != nil {
        return none, fmt.Errorf("Failed to url.JoinPath : %w", err)
    }

    if len(param) != 0 {
        endpoint += "?"
        first := true
        for k,v := range param {
            if !first {
                endpoint += "&"
            }
            first = false
            endpoint += k + "=" + v
        }
    }
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return none, fmt.Errorf("Failed to new request : %w", err)
	}
	req.SetBasicAuth(conf.UserName, conf.Password)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return none, fmt.Errorf("Failed to Do request : %w", err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

    var t T
    err =  json.Unmarshal(b, &t)

	return t, err
}

func toPath[T any](t T) string {
    if _, ok :=  any(t).(inkdrop.Notes) ; ok {
        return "notes"
    }
    if _, ok := any(t).(inkdrop.Note); ok {
        return "note"
    }
    if _, ok :=  any(t).(inkdrop.Books) ; ok {
        return "books"
    }
    if _, ok :=  any(t).(inkdrop.Tags) ; ok {
        return "tags"
    }

    return ""
}
