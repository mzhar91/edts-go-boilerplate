package security

import (
	"fmt"
	"regexp"
)

func CheckPasswordLever(ps string, length int, withSym bool) error {
	if len(ps) < length {
		return fmt.Errorf("password len is < %v", length)
	}
	
	num := `[0-9]{1}`
	az := `[a-z]{1}`
	AZ := `[A-Z]{1}`
	
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		// return fmt.Errorf("Password need num :%v", err)
		return fmt.Errorf("password need at least 1 number character")
	}
	
	if b, err := regexp.MatchString(az, ps); !b || err != nil {
		// return fmt.Errorf("Password need lowercase :%v", err)
		return fmt.Errorf("password need at least 1 lowercase character")
	}
	
	if b, err := regexp.MatchString(AZ, ps); !b || err != nil {
		// return fmt.Errorf("Password need A_Z :%v", err)
		return fmt.Errorf("password need at least 1 uppercase character")
	}
	
	if withSym {
		symbol := `[!@#~$%^&*()+|_-]{1}`
		
		if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
			// return fmt.Errorf("Password need symbol :%v", err)
			return fmt.Errorf("password need at least 1 symbol")
		}
	}
	
	return nil
}
