package Customer

import ("testing"
Restaurant "hotel/Restaurant"
)
var accForTest *Customer=New("neha", "neha@123", 5612345892, "@124", 1000)



func TestPlaceOrder(t *testing.T){
	rajpurohit := Restaurant.New("purohit")
	rajpurohit.AddNewItem("pav bhaji", 80, true, 4)
	accForTest.PlaceOrder(rajpurohit, "pav bhaji")
	var expected uint32=900
	actual:=accForTest.GetBalance()
	if expected!=actual{

		t.Errorf("expected was %v,but actual is %v",expected,actual)
	}

	
	
}
