package bankcustomer
import ("testing"
"fmt")
var accForTestsender *customer=New(1,"neha",334,5000,"savings",12)
var accForTestreciver *customer=New(2,"pooja",384,2000,"savings",12)
func TestMoneyTranfer(t *testing.T){
	
	err:=accForTestsender.MoneyTransfer(accForTestreciver, 334, 384, 12, 1000, "savings")
	if err!=nil{
		fmt.Println("there is an error ")
	}else{
		expected:=uint64(4000)
		t1,_:=accForTestsender.GetAccFromNo(334)
		actual:=t1.GetBalance()
		if expected!=actual{
			t.Errorf("expected %v,actual %v",expected,actual)
		}
	}
	
}
