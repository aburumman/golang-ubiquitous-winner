package camell

import "testing"


func TestIsCamell(t *testing.T) {
	if IsCamellEmployee("michael@camell.com",) == false {
		t.Errorf("Expected false but got true")
	}
}


func TestIsCamell2(t *testing.T) {
	if IsCamellEmployee("michael@google.com",) == true {
		t.Errorf("Expected false but got true")
	}
}

func TestIsCamellEmployeeTable(t *testing.T) {
	testCases := []struct {
		inputEmail string 
		ExpectedOutput bool
	}{
		{
			"michael@google.com",
			true
		},
		{
			"michael@microsoft.com",
			true
		},
		{
			"michael@linkedin.com",
			true
		}
	}
	}

for _, testCase := range testCases {
	t.Run(testCase.inputEmail, func(t *testing.T){
		actualOutput := IsCamellEmployee(testCase.inputEmail)
	})