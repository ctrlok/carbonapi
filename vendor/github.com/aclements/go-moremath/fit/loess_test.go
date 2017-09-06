// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fit

import (
	"testing"

	"github.com/aclements/go-moremath/internal/mathtest"
)

func TestLOESS_NIST(t *testing.T) {
	// LOWESS example from the NIST handbook.
	xs := []float64{0.5578196,
		2.0217271,
		2.5773252,
		3.4140288,
		4.3014084,
		4.7448394,
		5.1073781,
		6.5411662,
		6.7216176,
		7.2600583,
		8.1335874,
		9.1224379,
		11.9296663,
		12.3797674,
		13.2728619,
		14.2767453,
		15.3731026,
		15.6476637,
		18.5605355,
		18.5866354,
		18.7572812,
	}
	ys := []float64{18.63654,
		103.49646,
		150.35391,
		190.51031,
		208.70115,
		213.71135,
		228.49353,
		233.55387,
		234.55054,
		223.89225,
		227.68339,
		223.91982,
		168.01999,
		164.95750,
		152.61107,
		160.78742,
		168.55567,
		152.42658,
		221.70702,
		222.69040,
		243.18828,
	}

	defer mathtest.SetAeqDigits(mathtest.SetAeqDigits(7))
	mathtest.WantFunc(t, "LOESS", LOESS(xs, ys, 1, 0.33),
		map[float64]float64{
			0.5578196: 20.59302,
			2.0217271: 107.1603,
			2.5773252: 139.7674,
			3.4140288: 174.2630,
			4.301408:  207.2334,
			4.744839:  216.6616,
			5.107378:  220.5445,
			6.541166:  229.8607,
			6.721618:  229.8347,
			7.260058:  229.4301,
			8.133587:  226.6045,
			9.122438:  220.3904,
			11.929666: 172.3480,
			12.379767: 163.8417,
			13.272862: 161.8490,
			14.27675:  160.3351,
			15.37310:  160.1920,
			15.64766:  161.0556,
			18.56054:  227.3400,
			18.58664:  227.8985,
			18.75728:  231.5586,
		})
}