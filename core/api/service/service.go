package service

//func LoginService(c *gin.Context) {
//	var req protocol.Login
//	if err := c.ShouldBindJSON(&req); err != nil {
//		response.Response(http.StatusBadRequest, err.Error(), "", c)
//		return
//	}
//	err := req.Validate()
//	if err != nil {
//		response.Response(http.StatusBadRequest, err.Error(), "", c)
//		return
//	}
//	log.Debug(req)
//	code, msg, data := handler.AppLogin(config.JwtConfig().Exp, config.JwtConfig().Iss, config.JwtConfig().Sub, req.Username, req.Email)
//	response.Response(code, msg, data, c)
//	return
//}
//
//// DataUserRegisterService 数权用户注册
//func DataUserRegisterService(c *gin.Context) {
//	username := c.Request.URL.Query().Get("userName")
//	password := c.Request.URL.Query().Get("password")
//	if username == "" || password == "" {
//		response.Response(http.StatusBadRequest, response.ErrParameter, "", c)
//		return
//	}
//	code, msg, data := handler.DRUserRegister(username, password)
//	response.Response(code, msg, data, c)
//	return
//}
//func UploadService(c *gin.Context) {
//	did := c.Request.FormValue("did")
//	//去数据库查询did是否存在
//	log.Debug(did)
//	file, err := c.FormFile("file")
//	if err != nil {
//		response.Response(http.StatusInternalServerError, err.Error(), "", c)
//		return
//	}
//	open, err := file.Open()
//	if err != nil {
//		response.Response(http.StatusInternalServerError, err.Error(), "", c)
//		return
//	}
//	x, err := io.ReadAll(open)
//	if err != nil {
//		response.Response(http.StatusInternalServerError, err.Error(), "", c)
//		return
//	}
//	//d
//	handler.Upload(did, file.Filename, x)
//
//	//fmt.Println(file)
//}
