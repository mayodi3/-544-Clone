package main

func GetMenu(state string) string {
	response := ""

	switch state {
	case "page1":
		response = `CON 0:GEN Z OFFERS
			1:Social Bundles
			2:Data B-Live (NEW)
			3:Sh10=1GB YouTube, 3hrs
			4:Sh10=1.2GB, 1Hr
			5:MakeUrHOOK
			6:Sh20=200MB, 24Hr
			7:My Dealz
			8:Daily + Unlimited Music & Videos
			98:MORE`
	case "page2":
		response = `CON 9:Weekly+FREE Youtube
			10:GO MONTHLY
			11:DEALS MOTO
			12:Funzone+Partners
			13:GIFT
			14:Minutes
			15:Okoa
			16:Lipa Mdogo
			18:Data Kredo (Extra Data)
			19:Sh250=1GB+50Mins+100SMS
			0:BACK 98:MORE`
	case "page3":
		response = `CON 20:Sh500=7.5GB,30Days
			21:Bal & Tips
			22:No Expiry
			23:SMS Bundles
			24:Stop Auto Renew
			25:Unsubscribe From S-HOOK
			0:BACK`
	case "pageGenZ":
		response = `CON 1:Ksh20=1GB + (500GB) 1Hr
		2:Sh10=1GB 1Hr TikTok
		3:Sh20=200MB, 24hr + Unlimited Music and Videos
		4.Sh50=1.5GB, 3hr+Unlimited Music and Videos
		0:Back 00:Home 98:MORE`
	case "pageGenZNext":
		response = `CON 5:Sh 100=2GB, 24hr
		6:Night Hook Bundle (sh50=5GB midnight-6am)
		7:Sh10=15min, 1Hr
		0:Back 00:Home`
	case "socialBundles":
		response = `CON 1:TikTok
		2:YouTube
		3:Snapchat
		0:BACK 00:HOME`
	case "dataBlive":
		response = `CON Browse Worry-Free with NEW B-LIVE data.
		1:Sh30=Browse for 90Mins
		2:Sh75=Browse for 4.5Hours`
	case "dataBlivePage1":
		response = `CON Buy Sh30=Browse for 90Mins
		Pay With:
		1:Airtime
		2:Mpesa
		3:Pay with Okoa`
	case "dataBlivePage2":
		response = `CON Buy Sh75=Browse 4.5Hours
		Pay With:
		1:Airtime
		2:Mpesa
		3:Pay with Okoa`
	case "youtube1GB":
		response = `CON Buy 1GB YouTube Hourly @Ksh 10
		Pay via:
		1:Airtime
		2:Mpesa
		0:BACK 00:HOME`
	default:
		response = "END No Such Menu"
	}

	return response
}
