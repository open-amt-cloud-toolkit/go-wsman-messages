/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package common

import (
	"testing"
)

func TestDecodeWSMANError(t *testing.T) {
	tests := []struct {
		input    ErrorResponse
		expected error
	}{
		{
			ErrorResponse{
				Body: ErrorBody{
					Fault: Fault{
						Code: Code{
							SubCode: SubCode{
								Value: "b:AccessDenied",
							},
						},
						Reason: Reason{
							Text: "The sender was not authorized to access the resource.",
						},
						Detail: "",
					},
				},
			},
			&AMTError{
				SubCode: "b:AccessDenied", Message: "The sender was not authorized to access the resource.", Detail: "",
			},
		}, {
			ErrorResponse{
				Body: ErrorBody{
					Fault: Fault{
						Code: Code{
							SubCode: SubCode{
								Value: "e:DestinationUnreachable",
							},
						},
						Reason: Reason{
							Text: "No route can be determined to reach the destination role defined by the WSAddressing To.",
						},
						Detail: "",
					},
				},
			},
			&AMTError{
				SubCode: "e:DestinationUnreachable", Message: "No route can be determined to reach the destination role defined by the WSAddressing To.", Detail: "",
			},
		}, {
			ErrorResponse{},
			nil,
		},
	}

	for _, test := range tests {
		result := DecodeAMTError(test.input)
		if result == nil {
			if test.expected != nil {
				t.Errorf("Expected %s, but got nil", test.expected)
			}
		} else {
			if result.Error() != test.expected.Error() {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		}
	}
}
