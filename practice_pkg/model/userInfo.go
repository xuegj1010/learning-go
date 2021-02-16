package model

type UserInfo struct {
	Name      string
	Age       int
	Height    float32
	EduSchool string
	Hobby     []string
	MoreInfo  map[string]interface{}
}
