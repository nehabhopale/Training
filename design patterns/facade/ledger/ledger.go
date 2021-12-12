package ledger

import "fmt"

type Ledger struct {
}

func (s *Ledger) MakeEntry(accNo uint64, txt string, amount int) {
	fmt.Println(" entry for accountnumber", accNo, "with txt", txt, "for amount ", amount)

}
