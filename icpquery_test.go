package icpquery

import (
	"log"
	"testing"
)

func TestWebquery(t *testing.T) {
	icp, err := ICPQuery("qq.com")
	if err != nil {
		t.Fatal(err)
	}
	log.Print(icp)

	icp2, err := ICPQueryAizhan("qq.com")
	if err != nil {
		t.Fatal(err)
	}
	log.Print(icp2)
}
