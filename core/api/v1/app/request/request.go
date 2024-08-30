package request

type AppLoginRequest struct {
	AppName     string `json:"app_name"`
	AppAccount  string `json:"app_account"`
	AppPassword string `json:"app_password"`
}

type AppRegisterRequest struct {
	AppName     string `json:"app_name"`
	AppUser     string `json:"app_user"`
	AppPassword string `json:"app_password"`
	AppPhone    string `json:"app_phone"`
	AppType     string `json:"app_type"`
	Algo        string `json:"algo"`
	StoreKey    string `json:"store_key"`
}

type AppUserRegisterRequest struct {
	//加上AppName 防止篡token与业务应用,在Service层进行校验
	AppName        string `json:"app_name"`
	UserName       string `json:"user_name"`
	NickName       string `json:"nick_name"`
	Password       string `json:"password"`
	UserLevel      string `json:"user_level"`
	UserPhone      string `json:"user_phone"`
	UserSourceType string `json:"user_source_type"`
	UserSource     string `json:"user_source"`
	UserType       string `json:"user_type"`
}
type AppTestRequest struct {
	Name string `json:"name"`
}

type singleCount struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type AppAccountBindRequest struct {
	AppName   string        `json:"app_name"`
	DID       string        `json:"did"`
	AppCounts []singleCount `json:"app_counts"`
}

type AppInfoRequest struct {
	AppName  string `json:"app_name"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

type AppStatRequest struct {
	AppName string `json:"app_name"`
}
type AppUserCARequest struct {
	AppName string `json:"app_name"`
	DID     string `json:"did"`
}

type ServiceUploadRequest struct {
	AppName       string `json:"app_name" form:"app_name"`
	DID           string `json:"did" form:"did"`
	Password      string `json:"password" form:"password"`
	FileUrl       string `json:"file_url" form:"file_url"`
	FileName      string `json:"file_name" form:"file_name"`
	FileMD5       string `json:"file_md5" form:"file_md5"`
	FileRemark    string `json:"file_remark" form:"file_remark"`
	FileDirectory string `json:"file_directory" form:"file_directory"` //文件目录
	FileType      string `json:"file_type" form:"file_type"`           //1表示文档2表示图片3表示视频4表示关系数据
	CallbackUrl   string `json:"callback_url" form:"callback_url"`     //数据量太大的时候支持的http回调接口
}
type InternalDirectoryRequest struct {
	DirectID      int    `json:"direct_id"`
	DirectoryName string `json:"directory_name"`
	ParentID      int    `json:"parent_id"`
	Tag           string `json:"tag"`
	Remark        string `json:"remark"`
}

type ServiceDirectorySubscribeRequest struct {
	AppName string `json:"app_name"`
	RootDir int    `json:"root_dir"`
}

//	type ServiceUploadRequest struct {
//		AppName       string ` form:"app_name"`
//		DID           string `form:"did"`
//		Password      string ` form:"password"`
//		FileUrl       string `form:"file_url"`
//		FileName      string ` form:"file_name"`
//		FileMD5       string `form:"file_md5"`
//		FileRemark    string ` form:"file_remark"`
//		FileDirectory string ` form:"file_directory"` //文件目录
//		FileType      string ` form:"file_type"`      //1表示文档2表示图片3表示视频4表示关系数据
//		CallbackUrl   string ` form:"callback_url"`   //数据量太大的时候支持的http回调接口
//	}
type ServiceBuyRequest struct {
	AppName   string ` form:"app_name"`
	ServiceID int    ` json:"service_id"`
	//ServiceType string ` json:"service_type"`
	ServiceExp int64  ` json:"service_exp"`
	UseType    string ` json:"use_type"`
}
type ServiceListRequest struct {
	AppName  string ` json:"app_name"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

type ServiceDataListRequest struct {
	AppName    string `json:"app_name"`
	PageNum    int    `json:"page_num"`
	PageSize   int    `json:"page_size"`
	Conditions string `json:"conditions"` //条件,待确定
	NodeID     string `json:"node_id"`
	Directory  string `json:"directory"`
}

type ServiceWaterAddRequest struct {
	AppName     string `json:"app_name"`
	ServiceName string `json:"service_name"`
	Action      int    `json:"action"`
	Type        string `json:"type"`
	BucketName  string `json:"bucket_name"`
	ObjectName  string `json:"object_name"`
	Parameters  struct {
		CompanyName string `json:"company_name"`
		Content     string `json:"content"`
		Font        string `json:"font"`
		FontSize    string `json:"font_size"`
		FontColor   string `json:"font_color"`
		Strength    string `json:"strength"`
		Array       string `json:"array"`
		Key         string `json:"key"`
		Width       string `json:"width"`
		Height      string `json:"height"`
		Alpha       string `json:"alpha"`
		HideFiles   []struct {
			BucketName string `json:"bucket_name"`
			ObjectName string `json:"object_name"` //文档ID
		} `json:"hide_files"` //需要隐藏的文档ID列表，为空表示不隐藏任何文档
	} `json:"parameters"`
}

type CgRequest struct {
	AppName   string `json:"app_name"`
	OwnerDID  string `json:"owner_did"`
	PeerDID   string `json:"peer_did"`
	Cg        string `json:"cg"`      //g(grant) 授权(用peerDID的公钥加密) c(confirm) 确权(用ownerDID的私钥前签名) cg 先公钥加密在签名
	CgData    string `json:"cg_data"` //待cg数据
	CgDataMD5 string `json:"cg_data_md5"`
}

type CryptoRequest struct {
	AppName     string `json:"app_name" form:"app_name"`
	DataAID     string `json:"data_aid" form:"app_name"`
	DataType    string `json:"data_type" form:"app_name"`
	DataUrl     string `json:"data_url" form:"app_name"`
	DID         string `json:"did" form:"app_name"`
	Algo        string `json:"algo" form:"algo"`
	Action      int    `json:"action" form:"action"`
	CallbackUrl string `json:"callback_url" form:"callback_url"`
}

type StoreRequest struct {
	AppName  string `json:"app_name" form:"app_name"`
	DID      string `json:"did" form:"app_name"`
	FileName string `json:"file_name" form:"file_name"`
	FileUrl  string `json:"file_url" form:"file_url"`
	FileMD5  string `json:"file_md5" form:"file_md5"`
}

type SpaceInfoRequest struct {
	AppName  string `json:"app_name"`
	DID      string `json:"did"`
	Password string `json:"password"` //用户的密码
}

type ExpandRequest struct {
	AppName         string  `json:"app_name"`
	DID             string  `json:"did"`
	Password        string  `json:"password"` //用户的密码
	ExpandSpaceSize int     `json:"expand_space_size"`
	Price           float64 `json:"price"`
}

type UpdateRequest struct {
	AppName  string `json:"app_name"`
	DID      string `json:"did"`
	Password string `json:"password"` //
	DataID   string `json:"data_id"`
	DataName string `json:"data_name"`
	DataType string `json:"data_type"`
	IsOpen   string `json:"is_open"`
	DataPos  string `json:"data_pos"`
	NodeID   string `json:"node_id"`
	Remark   string `json:"remark"`
}

type DeleteRequest struct {
	AppName  string `json:"app_name"`
	Password string `json:"password"`
	DID      string `json:"did"`
	DataID   string `json:"data_id"`
}
type ShareRequest struct {
	AppNameSeller     string `json:"app_name_seller"`
	AppNameBuyer      string `json:"app_name_buyer "`
	DIDSeller         string `json:"did_owner"`
	DIDBuyer          string `json:"did_peer"`
	DataID            string `json:"data_id"`
	BuyerCallbackUrl  string `json:"buyer_callback"`
	SellerCallbackUrl string `json:"seller_callback"`
}

type RulesRequest struct {
	AppName   string `json:"app_name"`
	DataID    string `json:"data_id"`
	DID       string `json:"did"`
	SafeLevel string `json:"safe_level"`
	SafeInfo  string `json:"safe_info"`
}
