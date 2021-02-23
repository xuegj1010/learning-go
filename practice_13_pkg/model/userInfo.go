package model

type UserInfo struct {
	Name      string
	Age       int
	Height    float32
	EduSchool string
	Hobby     []string
	MoreInfo  map[string]interface{}
}
type userInfo1 struct {
	Name      string
	Age       int
	Height    float32
	EduSchool string
	Hobby     []string
	MoreInfo  map[string]interface{}
}

//工厂模式：生成对象
func NewUserInfo(name string, age int, height float32, edushool string, hobby []string, moreinfo map[string]interface{}) *userInfo1 {
	return &userInfo1{
		Name:      name,
		Age:       age,
		Height:    height,
		EduSchool: edushool,
		Hobby:     hobby,
		MoreInfo:  moreinfo,
	}
}
