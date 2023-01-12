package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muyisz/PracticalTraining/handler/body"
	"github.com/muyisz/PracticalTraining/internal"
	"github.com/muyisz/PracticalTraining/util"
)

func GetProductDetail(c *gin.Context) {
	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[GetProductDetail] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
	}
	reqBody := body.GetProductReq{}
	if err := util.PostForm(parms, &reqBody); err != nil {
		log.Printf("[GetProductDetail] BindJSON err,reqBody:%+v,err:%+v", reqBody, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}

	res, err := internal.GetProductDetail(reqBody.Pid)
	if err != nil {
		log.Printf("[GetProductDetail] GetProductDetail err,pid:%+v,err:%+v", reqBody.Pid, err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
		})
		return
	}

	c.JSON(http.StatusOK, body.Res{
		Code: 0,
		Data: res,
	})

}
