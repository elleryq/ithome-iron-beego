package controllers

// AboutController operations for About
type AboutController struct {
	BaseController
}

// URLMapping ...
func (c *AboutController) URLMapping() {
	c.Mapping("Get", c.Get)
}

func (c *AboutController) Get() {
	key := "asta"
	v := c.GetSession(key)
	if v == nil {
		c.SetSession(key, int(1))
		c.Data[key] = 0
	} else {
		c.SetSession(key, v.(int)+1)
		c.Data[key] = v.(int)
	}

	c.Data["Name"] = "John Doe"
	c.Data["Message"] = c.Tr("Cat is on the piano")
	c.TplName = "about.tpl"
}
