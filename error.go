// Copyright 2025- The sacloud/addon-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package addon

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-faster/errors"
	"github.com/sacloud/saclient-go"
)

type Error struct {
	msg string
	err error
}

type apiErrorResponse struct{ json.Marshaler }

func (e *Error) Unwrap() error { return e.err }
func (e *Error) Error() string {
	var buf strings.Builder

	buf.WriteString("addon")

	if e.msg != "" {
		buf.WriteString(": ")
		buf.WriteString(e.msg)
	}

	if e.err != nil {
		buf.WriteString(": ")
		buf.WriteString(e.err.Error())
	}

	return buf.String()
}

func (e *apiErrorResponse) Error() string {
	if buf, err := e.Marshaler.MarshalJSON(); err != nil {
		o := errors.Errorf("%#+v", e.Marshaler)
		return errors.Join(o, err).Error()
	} else {
		return string(buf)
	}
}

func NewError(msg string, err error) *Error { return &Error{msg: msg, err: err} }
func NewAPIError(method string, j json.Marshaler) *Error {
	return NewError(method, &apiErrorResponse{Marshaler: j})
}
func NewNotFoundError(method string) *Error {
	return NewError(method, saclient.NewError(http.StatusNotFound, "", nil))
}
