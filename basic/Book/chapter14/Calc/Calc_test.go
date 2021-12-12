package Calc
import "testing"

func TestAdd(t *testing.T){
	expected:=10
	actual:=Add(5,5)
	if expected!=actual{
		t.Errorf("expected %v,actual %v",expected,actual)
	}
}
func TestAddBulk(t *testing.T){
	var list=[]struct{
		num1,num2 int
		expected int
	}{
		{
			10,20,30,
		},{
			20,30,50,
		},
	}
	for _,val:=range list{
		actual :=Add(val.num1,val.num2)
		if val.expected!=actual{
			t.Errorf("exp was %v,actual is %v",val.expected,actual)
		}
	}
}
func TestSubBulk(t *testing.T){
	var list=[]struct{
		num1,num2 int
		expected int
	}{
		{
			30,20,10,
		},{
			23,3,20,
		},
	}
	for _,val:=range list{
		actual :=Sub(val.num1,val.num2)
		if val.expected!=actual{
			t.Errorf("exp was %v,actual is %v",val.expected,actual)
		}
	}
}
