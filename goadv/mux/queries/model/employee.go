package model

type Employee struct {
	
	COUNT int `gorm:"column:COUNT"`
	SUM float64 `gorm:"column:SUM"`
	AVG float64 `gorm:"column:AVG"`
	DEPTNO int `gorm:"column:DEPTNO"`
	JOB string `gorm:"column:JOB"`
	EMPNAME string `gorm:"column:EMPNAME"`
	DEPNAME string `gorm:"column:DEPNAME"`
	BOSSNAME string `gorm:"column:BOSSNAME"`
	RID int `gorm:"column:RID"`
	RNAME string `gorm:"column:RNAME"`
	CID string `gorm:"column:CID"`
	CNAME string `gorm:"column:CNAME"`
	STATEP string `gorm:"column:STATEP"`
	LID  int  `gorm:"column:LID"`
	SADD string `gorm:"column:SADD"`
	PCODE string `gorm:"column:PCODE"`
	CITY string `gorm:"column:CITY"`

}