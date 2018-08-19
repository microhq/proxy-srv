package handler

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/micro/go-micro/errors"
	proto "github.com/microhq/proxy-srv/proto"
	"golang.org/x/net/context"
)

type Proxy struct {
	Name string
	Url  string
}

func (p *Proxy) Call(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	// create url
	uri, err := url.Parse(p.Url)
	if err != nil {
		return errors.InternalServerError(p.Name+".proxy.url", err.Error())
	}
	uri.Path = req.Path

	// if we have get params make a get request
	var body io.Reader

	// if we have post params make a post request
	switch {
	// get request
	case len(req.Get) > 0:
		vals := url.Values{}
		for k, v := range req.Get {
			for _, val := range v.Values {
				vals.Add(k, val)
			}
		}
		body = bytes.NewReader([]byte(vals.Encode()))
	// post request
	case len(req.Post) > 0:
		vals := url.Values{}
		for k, v := range req.Post {
			for _, val := range v.Values {
				vals.Add(k, val)
			}
		}
		body = bytes.NewReader([]byte(vals.Encode()))
	// default body
	default:
		// maybe not efficient use of memory
		body = bytes.NewReader([]byte(req.Body))
	}

	// create the request
	r, err := http.NewRequest(req.Method, uri.String(), body)
	if err != nil {
		return errors.InternalServerError(p.Name+".proxy.call", err.Error())
	}

	// set the header
	for k, v := range req.Header {
		r.Header.Set(k, strings.Join(v.Values, ","))
	}

	// make the request
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return errors.InternalServerError(p.Name+".proxy.call.http", err.Error())
	}
	defer resp.Body.Close()

	// read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.InternalServerError(p.Name+".proxy.call.read-rsp-body", err.Error())
	}

	// set response
	rsp.Body = string(b)
	rsp.StatusCode = int32(resp.StatusCode)
	rsp.Header = make(map[string]*proto.Pair)

	for k, v := range resp.Header {
		rsp.Header[k] = &proto.Pair{
			Key:    k,
			Values: v,
		}
	}

	// we're done
	return nil
}
