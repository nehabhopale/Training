package account
import "fmt"

type SmsSub struct{
	subName string
	
}

func (s SmsSub) BalanceModified(a Account){
	fmt.Println("balance is modified for user",a.accUserName,"having subscription",s.subName)
}
