/**
 ******************************************************************************
 * @file           : pub.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"nunu/backend/IService"
	"nunu/backend/api/v1/base"
	"nunu/backend/dto"
)

type PubApi struct {
	Pub IService.IPubService
}

func NewPubAdmin(pub IService.IPubService) *PubApi {
	return &PubApi{Pub: pub}
}

func (p *PubApi) SendCaptchaWithMailReq(c *gin.Context) {
	req := new(dto.SendCaptchaWithMailReq)
	if err := base.Bind(c, req); err != nil {
		base.Render(c, nil, err)
		return
	}
	err := p.Pub.SendCaptchaWithMailReq(req)
	base.Render(c, nil, err)
}

func (p *PubApi) Register(c *gin.Context) {
	req := new(dto.RegisterReq)
	if err := base.Bind(c, req); err != nil {
		base.Render(c, nil, err)
		return
	}
	resp, err := p.Pub.Register(req)
	base.Render(c, resp, err)
}

func (p *PubApi) GetCaptcha(c *gin.Context) {
	resp, err := p.Pub.GetCaptcha()
	base.Render(c, resp, err)
}

func (p *PubApi) LoginWithUserName(c *gin.Context) {
	req := new(dto.LoginReqWithUserName)
	if err := base.Bind(c, req); err != nil {
		base.Render(c, nil, err)
		return
	}
	resp, err := p.Pub.LoginWithUsername(req)
	base.Render(c, resp, err)
}
