package securityCode

type SecurityCode struct {
	code uint16
}

func (c SecurityCode) CheckCode(code uint16) bool {
	return c.code == code

}
func New(code uint16) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}
