// Code generated by goctl. DO NOT EDIT.
package types

type CategoryCreateReq struct {
	NameZh string `json:"nameZh"`
	NameEn string `json:"nameEn"`
	Star   int64  `json:"star"`
}

type CategoryResp struct {
	Id        int64  `json:"id"`
	OwnerId   int64  `json:"ownerId"`
	NameZh    string `json:"nameZh"`
	NameEn    string `json:"nameEn"`
	Star      int64  `json:"star"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CategoryUpdateReq struct {
	PathIdReq
	CategoryCreateReq
}

type CodeSendReq struct {
	Mobile string `json:"mobile"`
}

type CodeValidateReq struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

type FormIdReq struct {
	Id int64 `form:"id"`
}

type GroupContactAddReq struct {
	GroupIdReq
	UserIds []int64 `json:"userIds"`
}

type GroupContactCategorySetReq struct {
	PathIdReq
	CategoryId int64 `json:"categoryId"`
}

type GroupContactIsSet struct {
	PathIdReq
	Value bool `json:"value"`
}

type GroupContactResp struct {
	Id             int64  `json:"id"`
	GroupId        int64  `json:"groupId"`
	UserId         int64  `json:"userId"`
	CategoryId     int64  `json:"categoryId"`
	UserNickname   string `json:"userNickname"`
	Remark         string `json:"remark"`
	Background     string `json:"background"`
	IsDisturb      bool   `json:"isDisturb"`
	IsTop          bool   `json:"isTop"`
	IsShowNickname bool   `json:"isShowNickname"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Category       struct {
		Id        int64  `json:"id"`
		OwnerId   int64  `json:"ownerId"`
		NameZh    string `json:"nameZh"`
		NameEn    string `json:"nameEn"`
		Star      int64  `json:"star"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"category"`
	Group struct {
		Id        int64  `json:"id"`
		CustomId  string `json:"customId"`
		Name      string `json:"name"`
		OwnerId   int64  `json:"ownerId"`
		Avatar    string `json:"avatar"`
		Memo      string `json:"memo"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"group"`
	User struct {
		Id          int64  `json:"id"`
		CustomId    string `json:"customId"`
		Mobile      string `json:"mobile"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Sex         int64  `json:"sex"`
		Memo        string `json:"memo"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		LastLoginAt string `json:"lastLoginAt"`
	} `json:"user"`
}

type GroupContactStringSetReq struct {
	PathIdReq
	Value string `json:"value"`
}

type GroupCreateReq struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Memo   string `json:"memo"`
}

type GroupIdReq struct {
	GroupId int64 `path:"groupId"`
}

type GroupResp struct {
	Id        int64  `json:"id"`
	CustomId  string `json:"customId"`
	Name      string `json:"name"`
	OwnerId   int64  `json:"ownerId"`
	Avatar    string `json:"avatar"`
	Memo      string `json:"memo"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GroupUpdateReq struct {
	PathIdReq
	GroupCreateReq
}

type PasswordResp struct {
	IsSetPassword bool `json:"isSetPassword"`
}

type PathIdReq struct {
	Id int64 `path:"id"`
}

type PathIdStrReq struct {
	Id string `path:"id"`
}

type ResourceCreateReq struct {
	Src    string `json:"src"`
	Type   int64  `json:"type"`
	NameZh string `json:"nameZh"`
	NameEn string `json:"nameEn"`
}

type ResourceResp struct {
	Id        int64  `json:"id"`
	Src       string `json:"src"`
	Type      int64  `json:"type"`
	NameZh    string `json:"nameZh"`
	NameEn    string `json:"nameEn"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ResourceType struct {
	Type int64 `path:"type"`
}

type ResourceUpdateReq struct {
	PathIdReq
	ResourceCreateReq
}

type UploadReq struct {
	File string `form:"file"`
}

type UploadResp struct {
	Url string `json:"url"`
}

type UserContactBackgroundSetReq struct {
	PathIdReq
	Background string `json:"background"`
}

type UserContactCategorySetReq struct {
	PathIdReq
	CategoryId int64 `json:"categoryId"`
}

type UserContactIsSet struct {
	PathIdReq
	Value bool `json:"value"`
}

type UserContactResp struct {
	Id         int64  `json:"id"`
	OwnerId    int64  `json:"ownerId"`
	DstId      int64  `json:"dstId"`
	CategoryId int64  `json:"categoryId"`
	Background string `json:"background"`
	IsDisturb  bool   `json:"isDisturb"`
	IsTop      bool   `json:"isTop"`
	IsRemind   bool   `json:"isRemind"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	Category   struct {
		Id        int64  `json:"id"`
		OwnerId   int64  `json:"ownerId"`
		NameZh    string `json:"nameZh"`
		NameEn    string `json:"nameEn"`
		Star      int64  `json:"star"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"category"`
	User struct {
		Id          int64  `json:"id"`
		CustomId    string `json:"customId"`
		Mobile      string `json:"mobile"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Sex         int64  `json:"sex"`
		Memo        string `json:"memo"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		LastLoginAt string `json:"lastLoginAt"`
	} `json:"user"`
}

type UserLoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type UserPasswordReq struct {
	PathIdReq
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UserResp struct {
	Id          int64  `json:"id"`
	CustomId    string `json:"customId"`
	Mobile      string `json:"mobile"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Sex         int64  `json:"sex"`
	Memo        string `json:"memo"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	LastLoginAt string `json:"lastLoginAt"`
}

type UserUpdateReq struct {
	PathIdReq
	CustomId string `json:"customId"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Sex      int64  `json:"sex"`
	Memo     string `json:"memo"`
}
