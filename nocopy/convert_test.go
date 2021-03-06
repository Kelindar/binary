// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package nocopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert_Bools(t *testing.T) {
	v := Bools{true, false, true, true, false, false}

	b := boolsToBinary(&v)
	assert.NotEmpty(t, b)
	assert.Equal(t, []byte{0x1, 0x0, 0x1, 0x1, 0x0, 0x0}, b)

	o := binaryToBools(&b)
	assert.NotEmpty(t, b)
	assert.Equal(t, v, o)
}
