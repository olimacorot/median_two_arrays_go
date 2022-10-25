package median_two_arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {

    tests := map[string]struct {
        numsOne []int 
        numsTwo []int
        want float64
    } {
        "it_should_return_median_2_for_arrays_1_3_2": {
            numsOne: []int{1,3},
            numsTwo: []int{2},
            want: 2,
        },
        "it_should_return_median_2_5_for_arrays_1_2_3_4": {
            numsOne: []int{1,2},
            numsTwo: []int{3,4},
            want: 2.5,
        },
        "it_should_return_median_2_5_for_arrays_1_3_2_7": {
            numsOne: []int{1,3},
            numsTwo: []int{2,7},
            want: 2.5,
        },
        "it_should_return_median_5_for_arrays_1_4_6_7": {
            numsOne: []int{1,4},
            numsTwo: []int{6,7},
            want: 5,
        },
        "it_should_return_median_4_for_arrays_1_4_0_6_7": {
            numsOne: []int{1,4,0},
            numsTwo: []int{6,7},
            want: 4,
        },
    }

    for name, tt := range tests {
        t.Run(name, func(t *testing.T) {
            assert.Equal(t, tt.want, findMedianSortedArrays(tt.numsOne, tt.numsTwo))
        })
    }
}