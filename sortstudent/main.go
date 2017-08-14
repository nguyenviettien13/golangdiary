package main

import (
	"sort"
	"fmt"
)

type SinhVien struct {
	Name string
	Age  int
}

type SinhVienArray []SinhVien

func (s SinhVienArray) Len() int {
	return len(s)
}
func (s SinhVienArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SinhVienArray) Less(i, j int) bool {//truong hop i<j thi phai thoa man dieu kien gi
	//sap xep theo ten va tuoi
	if s[i].Name!=s[j].Name{
		return s[i].Name < s[j].Name
	}else  {
		return s[i].Age <s[j].Age
	}
}
func main()  {
	var sinhvien SinhVien
	sinhvien.Age=21
	sinhvien.Name="Nguyen Thi Hoan"

	var sinhvien1 SinhVien
	sinhvien1.Age=22
	sinhvien1.Name="Nguyen Viet Tien"
	 
	var sinhvien2 SinhVien
	sinhvien2.Name="Nguyen Ha Vu"
	sinhvien2.Age =23

	var sinhvien3 SinhVien
	sinhvien3.Name="Nguyen Minh Chau"
	sinhvien3.Age =23

	var sinhvien4 SinhVien
	sinhvien4.Name="Nguyen Duy Nhat"
	sinhvien4.Age =24

	var sinhvien5 SinhVien
	sinhvien5.Name="Nguyen Tien Danh"
	sinhvien5.Age =25

	var sinhvien7 SinhVien
	sinhvien7.Name="Nguyen Thanh Truong"
	sinhvien7.Age =23

	var sinhvien6 SinhVien
	sinhvien6.Name="Nguyen Ha Vu"
	sinhvien6.Age =24

	

	var sinhvienarray = SinhVienArray{sinhvien,sinhvien1,sinhvien2,sinhvien3,sinhvien4,sinhvien5, sinhvien6, sinhvien7}

	sort.Sort(sinhvienarray)
	for _ ,value := range sinhvienarray{
		fmt.Println(value.Name,value.Age)
	}

	//var array = []string{"abc","def","Au duong phong","Duong Danh Duc"}
	//sort.Strings(array)
	//for _ ,value := range array{
	//	fmt.Println(value)
	//}
}
