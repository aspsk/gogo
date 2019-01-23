package conzt

import "testing"

func TestConzt(t *testing.T) {

    if KB != kb { t.Error("KB != kb") }
    if MB != mb { t.Error("MB != mb") }
    if GB != gb { t.Error("GB != gb") }
    if TB != tb { t.Error("TB != tb") }
    if PB != pb { t.Error("PB != pb") }
    if EB != eb { t.Error("EB != eb") }
    if ZB != zb { t.Error("ZB != zb") }
    if YB != yb { t.Error("YB != yb") }

}
