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
func TestMulBulk(t *testing.T){
	var list=[]struct{
		num1,num2 int
		expected int
	}{
		{
			10,10,100,
		},{
			20,10,20,
		},
	}
	for _,val:=range list{
		actual :=Mul(val.num1,val.num2)
		if val.expected!=actual{
			t.Errorf("exp was %v,actual is %v",val.expected,actual)
		}
	}
}
func TestSqurtBulk(t *testing.T){
	var list=[]struct{
		num1 float64
		expected float64
	}{
		{
			25 ,5,
		},{
			49,7,
		},
	}
	for _,val:=range list{
		actual :=SquareRoot(val.num1)
		if val.expected!=actual{
			t.Errorf("exp was %v,actual is %v",val.expected,actual)
		}
	}
}
func TestPowBulk(t *testing.T){
	var list=[]struct{
		num1,pow float64
		expected float64
	}{
		{
			5.0,2.0,25.0,
		},{
			3.0,2.0,9.0,
		},
	}
	for _,val:=range list{
		actual :=Power(val.num1,val.pow)
		if val.expected!=actual{
			t.Errorf("exp was %v,actual is %v",val.expected,actual)
		}
	}
}