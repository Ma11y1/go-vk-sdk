
[![VK Developers](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://dev.vk.com/)

### Tool for working with VKontakte API

### **All requests have not been tested, so there may be errors in the return values, especially in JSON names**

# Install
```bash
go get github.com/Ma11y1/go-vk-sdk
```

# Example

### DirectAuth
```GO
package main

import (
	"context"
	"fmt"
	"go-vk-sdk/action"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
)

func main() {
	a := api.NewAPI()
	actions := action.NewRouter(a)

	requestDirectAuth := actions.Auth.UserDirectAuth().
		ClientID(111).
		ClientSecret("APP_SECRET").
		Scope(actor.ScopeUserMessages + actor.ScopeUserFriends).
		Username("+7username").
		Password("pwd")

	responseDirectAuth, err := requestDirectAuth.Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	if responseDirectAuth.Error != nil { // if vk server error auth
		fmt.Println(responseDirectAuth.Error.Description)
	}
    
	user := responseDirectAuth.User
} 
```

### CodeFlow
```GO
package main

import (
	"context"
	"fmt"
	"go-vk-sdk/action"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
)

func main() {
	a := api.NewAPI()
	actions := action.NewRouter(a)

	// Create URL for get auth code
	getAuthCodeURL, err := (&actor.UserCodeFlowAuthorizeURL{
		ClientID:    111,
		RedirectURI: actor.AuthRedirectURI,
		Scope:       actor.ScopeUserFriends,
	}).Build()
	if err != nil {
		panic(err)
	}

	fmt.Println(getAuthCodeURL)

	code, err := actor.NewCodeFlowAuthorizeCodeRawURL("https://redirect.domain#code=code12345")
	if err != nil {
		panic(err)
	}

	// Swap auth code to token
	response, err := actions.Auth.UserCodeFlow().
		ClientID(111).
		ClientSecret("APP_SECRET").
		Code(code.Code).
		RedirectURI(actor.AuthRedirectURI).
		Exec(context.Background())
	if err != nil {
		panic(err)
	}

	if response.Error.Type != "" {
		panic(response.Error.Error())
	}

	user := response.User
}
```

### Request
```GO
package main

import (
	"context"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/action"
)

func main() {
	a := api.NewAPI()
	actions := action.NewRouter(a)

	var user actor.User
	
	profileInfo, err := actions.Account.GetProfileInfo(&user).Exec(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println(profileInfo)
}
```

```GO
package main

import (
	"context"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/request"
)

func main() {
	a := api.NewAPI()
    
	var user actor.User
	
	// actions have not yet been completed
	requestGetProfileInfo := request.NewFriendsGetRequest(a, &user)

	responseGetProfileInfo, err := requestGetProfileInfo.Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	if responseGetProfileInfo.Error.Code != errors.None {
		fmt.Println(responseGetProfileInfo.Error.Message)
	} else {
		fmt.Println(responseGetProfileInfo.Response.Items) 
	}
}
```

### LongPoll User
```GO
package main

import (
	"context"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/longPollUser"
)

func main() {
	a := api.NewAPI()
    
	var user actor.User
	
	lp := longPollUser.NewLongPoll(a, &user, longPollUser.ExtraOptionsModeReceiveAttachments)

	for update := range lp.Updates() {
		switch update.Type {
		case longPollUser.EventTypeMessageNew:
			fmt.Println(update.Event.(longPollUser.EventMessageNew))
		}
	}
}
```
