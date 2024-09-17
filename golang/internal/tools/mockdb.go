package tools

import(
	"time"
)
type mockDB struct {}

var mockLoginDetails = map[string]loginDetails{
	"alex": {
		AuthToken: "1234aBD",
		username: "Alex",
	},
	"mark": {
		AuthToken: "456ABD",
		username: "mark",
	},
	"pick": {
		AuthToken: "789aBD",
		username: "pick",
	},
	"gekkie": {
		AuthToken: "1234aBD",
		username: "gekkie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 300,
		Username: "alex",
	},
	"pick": {
		Coins: 300,
		Username: "pick",
	},
	"mark": {
		Coins: 300,
		Username: "mark",
	},
	"gekkie": {
		Coins: 300,
		Username: "gekkie",
	},
}


func (database *mockDB) GetUserLoginDetails(username string) *loginDetails{
	//simulate db call
time.Sleep(time.Second * 1)
var clientData = loginDetails{}
clientData, ok := mockLoginDetails[username]
if !ok{
	return nil
}
return &clientData
}


func (database *mockDB) GetUserCoins(username string) *CoinDetails{
	//simulate db call
time.Sleep(time.Second * 1)
var clientData = CoinDetails{}
clientData, ok := mockCoinDetails[username]
if !ok{
	return nil
}
return &clientData
}

func (database *mockDB) setupDatabse() error {
	return nil
}