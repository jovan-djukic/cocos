// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package http

import "github.com/ultravioletrs/manager/manager"

type apiReq interface {
	validate() error
}

type createDomainReq struct {
	Pool   string `json:"pool"`
	Volume string `json:"volume"`
	Domain string `json:"domain"`
}

func (req createDomainReq) validate() error {
	if req.Pool == "" || req.Volume == "" || req.Domain == "" {
		return manager.ErrMalformedEntity
	}

	return nil
}
