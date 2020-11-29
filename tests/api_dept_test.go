package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/model"
	"testing"
)

func TestCreateDepartment(t *testing.T) {
	//427772502
	rsp, err := dingTalk.CreateDept("golang", 1)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestCreateDetailDept:创建详细的部门
func TestCreateDetailDept(t *testing.T) {

	dept := new(model.Dept)
	dept.Name = "测试部门"
	dept.ParentId = 1 //根部门
	dept.SourceIdentifier = "123"
	dept.Hide = true
	dept.ShareBalance = true
	dept.ParentBalanceFirst = true
	dept.DeptPermits = []int{411415721, 411415721, 412362849}
	dept.OuterDept = true
	dept.OuterDeptOnlySelf = false
	dept.UserPermits = []string{"manager164", "manager164"}
	dept.DeptPermits = []int{411415721, 411415721, 412362849}
	dept.OuterPermitUsers = []string{"manager164", "manager164"}
	dept.OuterPermitDepts = []int{411415721, 411415721, 412362849}
	dept.Order = 1
	dept.Extension = `{"职能":"人事"}`

	rsp, err := dingTalk.CreateDetailDept(dept)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetDeptDetail：获取部门详情
func TestGetDeptDetail(t *testing.T) {

	rsp, err := dingTalk.GetDeptDetail(428051259, "")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetDepartmentUserIds：获取部门用户userid列表
func TestGetDepartmentUserIds(t *testing.T) {

	rsp, err := dingTalk.GetDeptUserIds(1)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//GetDeptUserDetail：获取部门用户userid列表
func TestGetDeptUserDetail(t *testing.T) {

	rsp, err := dingTalk.GetDeptUserDetail(1, 0, 3, "")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}
