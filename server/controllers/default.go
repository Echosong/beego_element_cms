package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"kaiyu-web/dtos"
	"kaiyu-web/models"
	util "kaiyu-web/utils"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

//简单处理二级关系
func (c *MainController) stageCategory() {
	categoryDb := new(models.Category)
	q := dtos.CategoryQuery{}
	q.PageSize = 1000
	q.Type = 0
	list, _ := categoryDb.PageList(q)
	var pCategories []*models.Category
	var sonCategories []*models.Category
	for _, category := range list {
		if category.ParentId == 0 {
			category.Sort = 0
			for _, c2 := range list {
				if int64(c2.ParentId) == category.Id {
					category.Sort = +1
				}
			}
			if len(category.Banner) > 6 {
				var jsonObj []map[string]string
				err := json.Unmarshal([]byte(category.Banner), &jsonObj)
				if err != nil {
					return
				}
				category.Banner = jsonObj[0]["url"]
				if len(jsonObj) > 1 {
					category.Banner = jsonObj[0]["url"] + "," + jsonObj[1]["url"]
				}
			}
			pCategories = append(pCategories, category)
		} else {
			sonCategories = append(sonCategories, category)
		}
	}
	categories = list
	c.Data["pCate"] = pCategories
	c.Data["sonCate"] = sonCategories
}

var categories []*models.Category

// Prepare 初始化链接数据
func (c *MainController) Prepare() {
	configDb := new(models.Config)
	mp := configDb.GetConfigs("site")

	//友情链接
	adDb := new(models.Ad)
	q := dtos.AdQuery{}
	q.Type = 1
	ads, _ := adDb.PageList(q)
	c.Data["ads"] = ads

	//分类
	categoryDb := new(models.Category)
	cq := dtos.CategoryQuery{}
	cq.Type = 1
	categoryLink, _ := categoryDb.PageList(cq)
	c.Data["adCates"] = categoryLink

	c.Data["config"] = mp
	c.stageCategory()
}

func getCurrentCate(id int64) (c *models.Category) {
	for _, category := range categories {
		if category.Id == id {
			return category
		}
	}
	return nil
}

// Index 首页
// @router / [get]
func (c *MainController) Index() {
	cate := getCurrentCate(1)
	imgs := strings.Split(cate.Banner, ",")
	var imgMap []map[string]string
	for _, img := range imgs {
		cMap := map[string]string{}
		cMap["url"] = img
		imgMap = append(imgMap, cMap)
	}
	c.Data["imgs"] = imgMap
	c.Data["cate"] = getCurrentCate(1)
	m := new(models.Content)
	about, _ := m.GetByCode("about")

	c.Data["about"] = about
	c.Data["products"] = c.getProduct(0)
	partners := c.getProduct(24)

	var peers []map[string]string
	//处理多个图片
	for _, partner := range partners {
		var p []map[string]string
		json.Unmarshal([]byte(partner.Content), &p)
		peers = append(peers, p...)
	}
	if len(peers) > 12 {
		c.Data["partners"] = peers[0:12]
	} else {
		c.Data["partners"] = peers
	}
	//技术服务图片
	c.getService()

	//重要案例
	articleDb := new(models.Article)
	aquery := dtos.ArticleQuery{}
	aquery.CategoryId = 12
	articles, _ := articleDb.PageList(aquery)
	if len(articles) > 2 {
		c.Data["articles"] = articles[:3]
	}

	c.TplName = "index.html"
}

//获取自主产品
func (c *MainController) getProduct(categoryId int) (products []*models.Product) {
	var ids []int
	if categoryId == 0 {
		for _, category := range categories {
			if category.ParentId == 20 {
				ids = append(ids, int(category.Id))
			}
		}
	} else {
		ids = append(ids, categoryId)
	}
	m := new(models.Product)
	products, _ = m.ListByCategory(ids)
	return products
}

// About 关于我们
// @router /about [get]
func (c *MainController) About() {
	m := new(models.Content)
	about, _ := m.GetByCode("about")
	c.Data["about"] = about
	var mp []map[string]string
	json.Unmarshal([]byte(about.Bs), &mp)

	c.Data["bs"] = mp

	//主营业务
	articleDb := new(models.Article)
	aquery := dtos.ArticleQuery{}
	aquery.CategoryId = 16
	articles, _ := articleDb.PageList(aquery)
	c.Data["articles"] = articles

	c.TplName = "about.html"
}

// Contact 联系我们
// @router /contact [get]
func (c *MainController) Contact() {
	c.TplName = "contact.html"
}

// List 新闻列表
// @route /list [get]
func (c *MainController) List() {
	c.TplName = "list.html"
}

// Content 内容详情
// @router content/:id [get]
func (c *MainController) Content() {
	id, _ := strconv.Atoi(c.GetString(":id"))
	index, _ := c.GetInt("index")
	if index == 0 {
		articleDb := &models.Article{}
		one, _ := articleDb.GetById(int64(id))
		c.Data["one"] = one
	} else {
		productDb := new(models.Product)
		one, _ := productDb.GetById(int64(id))
		c.Data["one"] = one
	}
	c.TplName = "content.html"
}

// Product 产品
// @router /product [get]
func (c *MainController) Product() {
	c.Data["Title"] = "自主产品"
	c.Data["products"] = c.getProduct(0)
	c.TplName = "product.html"
}

// Search  搜索
// @router /search [get]
func (c *MainController) Search() {
	keyword := c.GetString("keyword")
	page, _ := c.GetInt("page")
	index, _ := c.GetInt("index")
	if page < 1 {
		page = 1
	}
	c.Data["keyword"] = keyword
	c.Data["index"] = index
	var count int
	if index == 0 {
		//文章搜索
		article := new(models.Article)
		articleQuery := dtos.ArticleQuery{}
		articleQuery.Title = keyword
		articleQuery.Page = page
		list, i := article.PageList(articleQuery)
		c.Data["data"] = list
		count = int(i)
	} else {
		// 产品搜索
		product := new(models.Product)
		productQuery := dtos.ProductQuery{}
		productQuery.Page = page
		productQuery.Title = keyword
		list, i := product.PageList(productQuery)
		c.Data["data"] = list
		count = int(i)
	}
	c.Data["count"] = count
	c.Data["pagebar"] = util.NewPager(page, count, 10,
		fmt.Sprintf("/search?keyword=%s&index=%d", keyword, index), true).ToString()
	c.TplName = "search.html"
}

// Quality 公司资质
// @router /quality [get]
func (c *MainController) Quality() {

	var sonCates []*models.Category
	for _, category := range categories {
		if category.ParentId == 7 {
			sonCates = append(sonCates, category)
		}
	}
	id, _ := strconv.Atoi(c.GetString("id"))
	c.Data["currentId"] = id
	c.Data["qcate"] = sonCates
	if id == 7 {
		id = 8
	}
	articleDb := new(models.Article)
	aquery := dtos.ArticleQuery{}
	aquery.CategoryId = id
	articles, _ := articleDb.PageList(aquery)
	c.Data["articles"] = articles

	c.TplName = "quality.html"
}

func (c *MainController) getService() {
	services := c.getProduct(17)
	if len(services) > 0 {
		contentMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(services[0].Content), &contentMap)
		if err != nil {
			return
		}
		//logs.Info(contentMap)
		c.Data["service"] = contentMap
	}
}

// Service @Tile 技术服务
// @router /service [get]
func (c *MainController) Service() {
	c.getService()
	c.TplName = "service.html"
}

// Partner @Title 合作伙伴
// @router /partner [get]
func (c *MainController) Partner() {
	ids := []int{24}
	m := new(models.Product)
	products, _ := m.ListByCategory(ids)
	var mp []map[string]interface{}
	for _, product := range products {
		p := make(map[string]interface{})
		p["Title"] = product.Title
		var imgs []map[string]string
		json.Unmarshal([]byte(product.Content), &imgs)
		p["Imgs"] = imgs
		mp = append(mp, p)
	}
	c.Data["partners"] = mp
	c.TplName = "partners.html"
}

// Case 客户案例
// @router /case [get]
func (c *MainController) Case() {
	//行业客户
	ids := []int{13}
	m := new(models.Product)
	products, _ := m.ListByCategory(ids)
	var imgs []map[string]string
	if len(products) > 0 {
		err := json.Unmarshal([]byte(products[0].Content), &imgs)
		if err != nil {
			return
		}
		c.Data["imgs"] = imgs
	}
	page, _ := c.GetInt("page")

	if page < 1 {
		page = 1
	}
	//主要案例
	articleDb := new(models.Article)
	aquery := dtos.ArticleQuery{}
	aquery.CategoryId = 12
	aquery.PageSize = 4
	aquery.Page = page
	articles, count := articleDb.PageList(aquery)
	c.Data["articles"] = articles
	c.Data["pagebar"] = util.NewPager(page, int(count), aquery.PageSize,
		fmt.Sprintf("/case"), true).ToString()
	c.TplName = "case.html"
}
