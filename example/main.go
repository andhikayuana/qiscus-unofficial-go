package main

import (
	"fmt"

	"github.com/andhikayuana/qiscus-unofficial-go/multichannel"
	"github.com/andhikayuana/qiscus-unofficial-go/sdk"
)

func main() {
	// Initiate client for Multichannel
	multichannelClient := multichannel.NewMultichannel("karm-gzu41e4e4dv9fu3f", "A5wmVwtjhD3CIiIlQVCA", "karm-gzu41e4e4dv9fu3f_admin@qismo.com")

	// Initiate client for Multichannel using creadential email and password admin
	// multichannelClient, err := multichannel.NewMultichannelFromCredential("kaqiscus@gmail.com", "garasipojok")
	// if err != nil {
	// 	panic(err)
	// }

	// Initiate client for Multichannel using environment variable
	// QISMO_APP_CODE, QISMO_ADMIN_TOKEN, QISMO_ADMIN_EMAIL and QISMO_BASE_URL
	// multichannelClient, err := multichannel.NewMultichannelFromEnv()
	// if err != nil {
	// 	panic(err)
	// }

	// Default Multichannel base is https://multichannel.qiscus.com, use SetAPIBase() to override.
	// multichannelClient.SetAPIBase("https://multichannel2.qiscus.com")

	// Sample Multichannel method
	resp, _ := multichannelClient.GetRoomTags("48627228")
	fmt.Println(resp)

	// Initiate client for SDK
	sdkClient := sdk.NewSDK("karm-gzu41e4e4dv9fu3f", "9d54f6697777234479341ed20dba7517")

	// Initiate client for SDK using environment variable
	// QISCUS_SDK_APP_CODE, QISCUS_SDK_SECRET_KEY and QISCUS_SDK_BASE_URL
	// sdkClient, err := sdk.NewSDKFromEnv()
	// if err != nil {
	// 	panic(err)
	// }

	// Default SDK base is https://api.qiscus.com, use SetAPIBase() to override.
	// sdkClient.SetAPIBase("https://api2.qiscus.com")

	// Sample SDK method
	resp2, _ := sdkClient.LoginOrRegister(&sdk.LoginOrRegisterReq{
		UserID:   "guest2@qiscus.com",
		Password: "123123123",
		Username: "User Demo",
	})
	fmt.Println(resp2)
}
