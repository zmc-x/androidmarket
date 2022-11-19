package manage

import (
	manageresp "androidmarket/model/common/response"
	"androidmarket/model/manage/request"
	"androidmarket/utils/upload"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Managegoods struct{}

// AddGoodsImages 添加商品图片
func (m Managegoods) AddGoodsImages(c *gin.Context) {
	// 获取需要添加图片的id
	goodsid, _ := strconv.Atoi(c.Param("goodsid"))
	// 获取图片信息
	form, err := c.MultipartForm()
	if err != nil {
		manageresp.FailParam(c)
		return
	}
	// 封面图
	coverimage := form.File["coverImage"]
	coversrc, covererr := upload.UploadFile(coverimage[0])
	if covererr != nil {
		manageresp.FailWithDetail(c, map[string]interface{}{"error": covererr.Error()}, "文件上传出错！")
		return
	}
	images := form.File["images"]
	// imagespath
	var imagespath string
	for _, v := range images {
		src, err := upload.UploadFile(v)
		if err != nil {
			manageresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, "文件上传出错！")
			return
		}
		imagespath = imagespath + src + ","
	}
	status, err := AddGoodsImage(coversrc, imagespath, goodsid)
	if status {
		manageresp.OkWithDetail(c, nil, "商品图片添加成功！")
	} else {
		if err != nil {
			manageresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, "商品图片添加失败！")
		} else {
			manageresp.FailWithDetail(c, nil, "商品图片添加失败！")
		}
	}
}

// AddGoodsMajorInfo 添加商品部分信息
func (m Managegoods) AddGoodsMajorInfo(c *gin.Context) {
	goodsinfo := request.Goodsinfo{}
	err := c.ShouldBindJSON(&goodsinfo)
	if err != nil {
		manageresp.FailParam(c)
		return
	}
	if status, err := AddGoodsMajorInfo(goodsinfo); status {
		manageresp.OkWithDetail(c, nil, "商品主要信息添加成功！")
	} else {
		if err != nil {
			manageresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, "商品主要信息添加失败！")
		} else {
			manageresp.FailWithDetail(c, nil, "商品主要信息添加失败！")
		}
	}
}

// AddGoodsSpecific 添加商品规格
func (m Managegoods) AddGoodsSpecific(c *gin.Context) {
	goodsinfo := request.Specification{}
	err := c.ShouldBindJSON(&goodsinfo)
	if err != nil {
		manageresp.FailParam(c)
		return
	}
	if msg, status, err := AddGoodsSpecific(goodsinfo); status {
		manageresp.OkWithDetail(c, nil, msg)
	} else {
		if err != nil {
			manageresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, msg)
		} else {
			manageresp.FailWithDetail(c, nil, msg)
		}
	}
}

// UpdateGoodsinfo 修改商品信息
func (m Managegoods) UpdateGoodsinfo(c *gin.Context) {
	goodsinfo := request.Updateinfo{}
	err := c.ShouldBindJSON(&goodsinfo)
	if err != nil {
		manageresp.FailParam(c)
		return
	}
	if msg, status, err := UpdateGoodsinfo(goodsinfo); status {
		manageresp.OkWithDetail(c, nil, msg)
	} else {
		if err != nil {
			manageresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, msg)
		} else {
			manageresp.FailWithDetail(c, nil, msg)
		}
	}
}

// DeleteGoodsinfo 删除商品 or 商品对应的规格
func (m Managegoods) DeleteGoodsinfo(c *gin.Context) {
	goodsinfo := request.Deleteinfo{}
	err := c.ShouldBindJSON(&goodsinfo)
	if err != nil {
		manageresp.FailParam(c)
		return
	}
	if msg, status := DeleteGoodsinfo(goodsinfo); status {
		manageresp.OkWithDetail(c, nil, msg)
	} else {
		manageresp.FailWithDetail(c, nil, msg)
	}
}
