package cryptocomparego

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSocialStatsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/socialstats", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("id") != "1182" {
			t.Errorf("SocialStats.Get did not request the correct id")
		}

		response := `
		{
		  "Response": "Success",
		  "Message": "Social data successfully returned",
		  "Type": 100,
		  "Data": {
			"General": {
			  "Name": "BTC",
			  "CoinName": "Bitcoin",
			  "Type": "Webpagecoinp",
			  "Points": 4149007
			},
			"CryptoCompare": {
			  "SimilarItems": [
				{
				  "Id": 7605,
				  "Name": "Ethereum",
				  "FullName": "Ethereum (ETH)",
				  "ImageUrl": "/media/20646/eth_logo.png",
				  "Url": "/coins/eth/",
				  "FollowingType": 1
				},
				{
				  "Id": 5031,
				  "Name": "Ripple",
				  "FullName": "Ripple (XRP)",
				  "ImageUrl": "/media/19972/ripple.png",
				  "Url": "/coins/xrp/",
				  "FollowingType": 1
				},
				{
				  "Id": 3808,
				  "Name": "Litecoin",
				  "FullName": "Litecoin (LTC)",
				  "ImageUrl": "/media/19782/litecoin-logo.png",
				  "Url": "/coins/ltc/",
				  "FollowingType": 1
				},
				{
				  "Id": 127356,
				  "Name": "IOTA",
				  "FullName": "IOTA (IOT)",
				  "ImageUrl": "/media/1383540/iota_logo.png",
				  "Url": "/coins/iot/",
				  "FollowingType": 1
				},
				{
				  "Id": 202330,
				  "Name": "Bitcoin Cash / BCC",
				  "FullName": "Bitcoin Cash / BCC (BCH)",
				  "ImageUrl": "/media/1383919/bch.jpg",
				  "Url": "/coins/bch/",
				  "FollowingType": 1
				},
				{
				  "Id": 5038,
				  "Name": "Monero",
				  "FullName": "Monero (XMR)",
				  "ImageUrl": "/media/19969/xmr.png",
				  "Url": "/coins/xmr/",
				  "FollowingType": 1
				},
				{
				  "Id": 4614,
				  "Name": "Stellar",
				  "FullName": "Stellar (XLM)",
				  "ImageUrl": "/media/20696/str.png",
				  "Url": "/coins/xlm/",
				  "FollowingType": 1
				},
				{
				  "Id": 27368,
				  "Name": "NEO",
				  "FullName": "NEO (NEO)",
				  "ImageUrl": "/media/1383858/neo.jpg",
				  "Url": "/coins/neo/",
				  "FollowingType": 1
				},
				{
				  "Id": 4433,
				  "Name": "Verge",
				  "FullName": "Verge (XVG)",
				  "ImageUrl": "/media/12318032/xvg.png",
				  "Url": "/coins/xvg/",
				  "FollowingType": 1
				},
				{
				  "Id": 24854,
				  "Name": "ZCash",
				  "FullName": "ZCash (ZEC)",
				  "ImageUrl": "/media/351360/zec.png",
				  "Url": "/coins/zec/",
				  "FollowingType": 1
				}
			  ],
			  "CryptopianFollowers": [
				{
				  "Id": 613700,
				  "Name": "sytze_vv",
				  "ImageUrl": "https://images.cryptocompare.com/613700/209646d5-b707-4b12-aef5-f649c3f85dc0.jpg",
				  "Url": "/profile/sytze_vv/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 237010,
				  "Name": "bruce.johnsen",
				  "ImageUrl": "https://images.cryptocompare.com/237010/759bc563-6c2b-4b78-a7b7-022630e6d4b5.jpg",
				  "Url": "/profile/bruce.johnsen/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 660353,
				  "Name": "Mario.sommer654",
				  "ImageUrl": "https://images.cryptocompare.com/660353/4871977a-7f86-4eb8-9e15-e5be2b1a82f0.jpg",
				  "Url": "/profile/Mario.sommer654/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 665164,
				  "Name": "pablo.altamirano",
				  "ImageUrl": "https://images.cryptocompare.com/665164/b631ef2b-658e-4c50-a05a-39faf50e0643.jpg",
				  "Url": "/profile/pablo.altamirano/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 415267,
				  "Name": "Rajshare",
				  "ImageUrl": "https://images.cryptocompare.com/415267/1a531fa8-fb70-4733-9c70-5ae5880c447e.jpg",
				  "Url": "/profile/Rajshare/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 665252,
				  "Name": "nagji9553",
				  "ImageUrl": "https://images.cryptocompare.com/665252/570aefc1-da0e-4729-910c-cf3a19866949.jpg",
				  "Url": "/profile/nagji9553/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 644974,
				  "Name": "60A68355",
				  "ImageUrl": "https://images.cryptocompare.com/644974/c5870b20-46ed-4a90-8c38-fa71eece5891.jpg",
				  "Url": "/profile/60A68355/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 185446,
				  "Name": "dfv219",
				  "ImageUrl": "https://images.cryptocompare.com/185446/2c2d07ee-369a-4caa-9d63-dec5f15de1ee.jpg",
				  "Url": "/profile/dfv219/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 665169,
				  "Name": "hkaynartr",
				  "ImageUrl": "https://images.cryptocompare.com/665169/a996b0f4-b537-4da8-947b-8bbaaee94f4d.jpg",
				  "Url": "/profile/hkaynartr/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 217556,
				  "Name": "lewis.john.h",
				  "ImageUrl": "https://images.cryptocompare.com/217556/acded0cf-ecba-425c-8445-3a5d4584a6da.jpg",
				  "Url": "/profile/lewis.john.h/",
				  "Type": "Cryptopian"
				},
				{
				  "Id": 664402,
				  "Name": "jvenjohn",
				  "ImageUrl": "https://images.cryptocompare.com/664402/6bdc9479-deb7-4448-b01a-bd42fe928bb3.jpg",
				  "Url": "/profile/jvenjohn/",
				  "Type": "Cryptopian"
				}
			  ],
			  "Points": 2626185,
			  "Followers": 45941,
			  "Posts": "32724",
			  "Comments": "65956",
			  "PageViewsSplit": {
				"Overview": 10815667,
				"Markets": 784901,
				"Analysis": 538435,
				"Charts": 3102916,
				"Trades": 326362,
				"Orderbook": 50035,
				"Forum": 1881598,
				"Influence": 33391
			  },
			  "PageViews": 17533305
			},
			"Twitter": {
			  "following": "114",
			  "account_creation": "1313643968",
			  "name": "Bitcoin",
			  "lists": 5745,
			  "statuses": 19229,
			  "favourites": "89",
			  "followers": 679451,
			  "link": "https://twitter.com/bitcoin",
			  "Points": 708195
			},
			"Reddit": {
			  "posts_per_hour": "14.58",
			  "comments_per_hour": "188.38",
			  "posts_per_day": "349.90",
			  "comments_per_day": 4521.19,
			  "name": "Bitcoin",
			  "link": "https://www.reddit.com/r/bitcoin/",
			  "active_users": 12191,
			  "community_creation": "1284042626",
			  "subscribers": 657642,
			  "Points": 703257
			},
			"Facebook": {
			  "likes": 36699,
			  "link": "https://www.facebook.com/bitcoins/",
			  "is_closed": "false",
			  "talking_about": "366",
			  "name": "Bitcoin P2P Cryptocurrency",
			  "Points": 36699
			},
			"CodeRepository": {
			  "List": [
				{
				  "created_at": "1363239994",
				  "open_total_issues": "25",
				  "parent": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "size": "1174",
				  "closed_total_issues": "948",
				  "stars": 699,
				  "last_update": "1515732848",
				  "forks": 236,
				  "url": "https://github.com/petertodd/python-bitcoinlib",
				  "closed_issues": "373",
				  "closed_pull_issues": "575",
				  "fork": "false",
				  "last_push": "1514847167",
				  "source": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "open_pull_issues": "10",
				  "language": "Python",
				  "subscribers": 91,
				  "open_issues": "24"
				},
				{
				  "created_at": "1304525025",
				  "open_total_issues": "34",
				  "parent": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "size": "2990",
				  "closed_total_issues": "139",
				  "stars": 2227,
				  "last_update": "1515739647",
				  "forks": 721,
				  "url": "https://github.com/bitcoinjs/bitcoinjs-lib",
				  "closed_issues": "46",
				  "closed_pull_issues": "655",
				  "fork": "false",
				  "last_push": "1515609223",
				  "source": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "open_pull_issues": "26",
				  "language": "JavaScript",
				  "subscribers": 149,
				  "open_issues": "14"
				},
				{
				  "created_at": "1384835603",
				  "open_total_issues": "221",
				  "parent": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "size": "16807",
				  "closed_total_issues": "11197",
				  "stars": 1863,
				  "last_update": "1515723880",
				  "forks": 1247,
				  "url": "https://github.com/bitcoinj/bitcoinj",
				  "closed_issues": "3040",
				  "closed_pull_issues": "93",
				  "fork": "false",
				  "last_push": "1515402478",
				  "source": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "open_pull_issues": "11",
				  "language": "Java",
				  "subscribers": 260,
				  "open_issues": "579"
				},
				{
				  "created_at": "1292771803",
				  "open_total_issues": "846",
				  "parent": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "size": "69261",
				  "closed_total_issues": "1272",
				  "stars": 25016,
				  "last_update": "1515750283",
				  "forks": 14581,
				  "url": "https://github.com/bitcoin/bitcoin",
				  "closed_issues": "617",
				  "closed_pull_issues": "8157",
				  "fork": "false",
				  "last_push": "1515743168",
				  "source": {
					"Name": "",
					"Url": "",
					"InternalId": -1
				  },
				  "open_pull_issues": "267",
				  "language": "C++",
				  "subscribers": 2681,
				  "open_issues": "195"
				}
			  ],
			  "Points": 72918
			}
		  }
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.SocialStats.Get(ctx, 1182)
	if err != nil {
		t.Errorf("SocialStats.Get returned error: %v", err)
	}

	expected := SocialStats{
		General: General{Name: "BTC", CoinName: "Bitcoin", Type: "Webpagecoinp", Points: 4149007},
		CryptoCompare: CryptoCompare{
			SimilarItems: []SimilarItem{
				{Id: 7605, Name: "Ethereum", FullName: "Ethereum (ETH)", ImageUrl: "/media/20646/eth_logo.png", Url: "/coins/eth/", FollowingType: 1},
				{Id: 5031, Name: "Ripple", FullName: "Ripple (XRP)", ImageUrl: "/media/19972/ripple.png", Url: "/coins/xrp/", FollowingType: 1},
				{Id: 3808, Name: "Litecoin", FullName: "Litecoin (LTC)", ImageUrl: "/media/19782/litecoin-logo.png", Url: "/coins/ltc/", FollowingType: 1},
				{Id: 127356, Name: "IOTA", FullName: "IOTA (IOT)", ImageUrl: "/media/1383540/iota_logo.png", Url: "/coins/iot/", FollowingType: 1},
				{Id: 202330, Name: "Bitcoin Cash / BCC", FullName: "Bitcoin Cash / BCC (BCH)", ImageUrl: "/media/1383919/bch.jpg", Url: "/coins/bch/", FollowingType: 1},
				{Id: 5038, Name: "Monero", FullName: "Monero (XMR)", ImageUrl: "/media/19969/xmr.png", Url: "/coins/xmr/", FollowingType: 1},
				{Id: 4614, Name: "Stellar", FullName: "Stellar (XLM)", ImageUrl: "/media/20696/str.png", Url: "/coins/xlm/", FollowingType: 1},
				{Id: 27368, Name: "NEO", FullName: "NEO (NEO)", ImageUrl: "/media/1383858/neo.jpg", Url: "/coins/neo/", FollowingType: 1},
				{Id: 4433, Name: "Verge", FullName: "Verge (XVG)", ImageUrl: "/media/12318032/xvg.png", Url: "/coins/xvg/", FollowingType: 1},
				{Id: 24854, Name: "ZCash", FullName: "ZCash (ZEC)", ImageUrl: "/media/351360/zec.png", Url: "/coins/zec/", FollowingType: 1},
			},
			CryptopianFollowers: []CryptopianFollower{
				{Id: 613700, Name: "sytze_vv", ImageUrl: "https://images.cryptocompare.com/613700/209646d5-b707-4b12-aef5-f649c3f85dc0.jpg", Url: "/profile/sytze_vv/", Type: "Cryptopian"},
				{Id: 237010, Name: "bruce.johnsen", ImageUrl: "https://images.cryptocompare.com/237010/759bc563-6c2b-4b78-a7b7-022630e6d4b5.jpg", Url: "/profile/bruce.johnsen/", Type: "Cryptopian"},
				{Id: 660353, Name: "Mario.sommer654", ImageUrl: "https://images.cryptocompare.com/660353/4871977a-7f86-4eb8-9e15-e5be2b1a82f0.jpg", Url: "/profile/Mario.sommer654/", Type: "Cryptopian"},
				{Id: 665164, Name: "pablo.altamirano", ImageUrl: "https://images.cryptocompare.com/665164/b631ef2b-658e-4c50-a05a-39faf50e0643.jpg", Url: "/profile/pablo.altamirano/", Type: "Cryptopian"},
				{Id: 415267, Name: "Rajshare", ImageUrl: "https://images.cryptocompare.com/415267/1a531fa8-fb70-4733-9c70-5ae5880c447e.jpg", Url: "/profile/Rajshare/", Type: "Cryptopian"},
				{Id: 665252, Name: "nagji9553", ImageUrl: "https://images.cryptocompare.com/665252/570aefc1-da0e-4729-910c-cf3a19866949.jpg", Url: "/profile/nagji9553/", Type: "Cryptopian"},
				{Id: 644974, Name: "60A68355", ImageUrl: "https://images.cryptocompare.com/644974/c5870b20-46ed-4a90-8c38-fa71eece5891.jpg", Url: "/profile/60A68355/", Type: "Cryptopian"},
				{Id: 185446, Name: "dfv219", ImageUrl: "https://images.cryptocompare.com/185446/2c2d07ee-369a-4caa-9d63-dec5f15de1ee.jpg", Url: "/profile/dfv219/", Type: "Cryptopian"},
				{Id: 665169, Name: "hkaynartr", ImageUrl: "https://images.cryptocompare.com/665169/a996b0f4-b537-4da8-947b-8bbaaee94f4d.jpg", Url: "/profile/hkaynartr/", Type: "Cryptopian"},
				{Id: 217556, Name: "lewis.john.h", ImageUrl: "https://images.cryptocompare.com/217556/acded0cf-ecba-425c-8445-3a5d4584a6da.jpg", Url: "/profile/lewis.john.h/", Type: "Cryptopian"},
				{Id: 664402, Name: "jvenjohn", ImageUrl: "https://images.cryptocompare.com/664402/6bdc9479-deb7-4448-b01a-bd42fe928bb3.jpg", Url: "/profile/jvenjohn/", Type: "Cryptopian"},
			},
			Comments: "65956", Points: 2626185, Posts: "32724", Followers: 45941, PageViewsSplit: PageViewsSplit{Overview: 10815667, Markets: 784901, Analysis: 538435, Charts: 3102916, Trades: 326362, Forum: 1881598, Influence: 33391}, PageViews: 17533305},
		Twitter:  Twitter{Followers: 679451, Following: "114", Lists: 5745, Favourites: "89", Statuses: 19229, AccountCreation: "1313643968", Name: "Bitcoin", Link: "https://twitter.com/bitcoin", Points: 708195},
		Reddit:   Reddit{Subscribers: 657642, ActiveUsers: 12191, CommunityCreation: "1284042626", PostsPerHour: 14.58, PostsPerDay: 349.9, CommentsPerHour: 188.38, CommentsPerDay: 4521.19, Link: "https://www.reddit.com/r/bitcoin/", Name: "Bitcoin", Points: 703257},
		Facebook: Facebook{Likes: 36699, IsClosed: false, TalkingAbout: "366", Name: "Bitcoin P2P Cryptocurrency", Link: "https://www.facebook.com/bitcoins/", Points: 36699},
		CodeRepository: CodeRepository{
			List: []CodeRepositoryItem{
				{Stars: 699, Language: "Python", Forks: 236, OpenTotalIssues: "25", Subscribers: 91, Size: "1174", Url: "https://github.com/petertodd/python-bitcoinlib", LastUpdate: "1515732848", LastPush: "1514847167", CreatedAt: "1363239994", Fork: false, Source: Source{Name: "", Url: "", InternalId: -1}, Parent: Source{Name: "", Url: "", InternalId: -1}, OpenPullIssues: "10", ClosedPullIssues: "575", ClosedTotalIssues: "948", OpenIssues: "24", ClosedIssues: "373"},
				{Stars: 2227, Language: "JavaScript", Forks: 721, OpenTotalIssues: "34", Subscribers: 149, Size: "2990", Url: "https://github.com/bitcoinjs/bitcoinjs-lib", LastUpdate: "1515739647", LastPush: "1515609223", CreatedAt: "1304525025", Fork: false, Source: Source{Name: "", Url: "", InternalId: -1}, Parent: Source{Name: "", Url: "", InternalId: -1}, OpenPullIssues: "26", ClosedPullIssues: "655", ClosedTotalIssues: "139", OpenIssues: "14", ClosedIssues: "46"},
				{Stars: 1863, Language: "Java", Forks: 1247, OpenTotalIssues: "221", Subscribers: 260, Size: "16807", Url: "https://github.com/bitcoinj/bitcoinj", LastUpdate: "1515723880", LastPush: "1515402478", CreatedAt: "1384835603", Fork: false, Source: Source{Name: "", Url: "", InternalId: -1}, Parent: Source{Name: "", Url: "", InternalId: -1}, OpenPullIssues: "11", ClosedPullIssues: "93", ClosedTotalIssues: "11197", OpenIssues: "579", ClosedIssues: "3040"},
				{Stars: 25016, Language: "C++", Forks: 14581, OpenTotalIssues: "846", Subscribers: 2681, Size: "69261", Url: "https://github.com/bitcoin/bitcoin", LastUpdate: "1515750283", LastPush: "1515743168", CreatedAt: "1292771803", Fork: false, Source: Source{Name: "", Url: "", InternalId: -1}, Parent: Source{Name: "", Url: "", InternalId: -1}, OpenPullIssues: "267", ClosedPullIssues: "8157", ClosedTotalIssues: "1272", OpenIssues: "195", ClosedIssues: "617"},
			},
			Points: 72918},
	}

	if !reflect.DeepEqual(*acct, expected) {
		t.Errorf("SocialStats.Get\nreturned\n%+v\nexpected\n%+v", *acct, expected)
	}
}
