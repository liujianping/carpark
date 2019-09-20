package prepare

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepare(t *testing.T) {
	csv, err := os.OpenFile("../dataset/hdb-carpark-information.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	assert.Nil(t, err)

	items, err := Parse(csv)
	assert.Nil(t, err)
	assert.Equal(t, 2113, len(items))

	d1, d2, err := parkingPeriod("7AM-10.30PM")
	assert.Nil(t, err)
	assert.Equal(t, 25200, int(d1.Seconds()))
	assert.Equal(t, 81000, int(d2.Seconds()))

	d3, d4, err := parkingPeriod("NO")
	assert.Nil(t, err)
	assert.Equal(t, 0, int(d3.Seconds()))
	assert.Equal(t, 0, int(d4.Seconds()))

	d5, d6, err := parkingPeriod("WHOLE DAY")
	assert.Nil(t, err)
	assert.Equal(t, 0, int(d5.Seconds()))
	assert.Equal(t, 86400, int(d6.Seconds()))

	_, _, err = parkingPeriod("7AM:10.30PM")
	assert.NotNil(t, err)
}
