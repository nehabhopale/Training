package accountdetails
import "testing"
var accForTest *Account=New(334,1,3000,"saving",6)
func TestGetBalance(t *testing.T){
	expected:=uint64(3000)

	actual:=accForTest.GetBalance()
	if expected!=actual{
		t.Errorf("expected %v,actual %v",expected,actual)
	}
}
func TestGetAccNo(t *testing.T){
	expected:=uint64(336)

	actual:=accForTest.GetAccNo()
	if expected!=actual{
		t.Errorf("expected %v,actual %v",expected,actual)
	}
}