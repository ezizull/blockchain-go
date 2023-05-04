package signature

import "fmt"

func (sig *Signature) String() string {
	return fmt.Sprintf("%x%x", sig.R, sig.S)
}
