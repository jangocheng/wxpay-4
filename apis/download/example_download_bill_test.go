// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package download_test

import (
	"fmt"

	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/apis/download"
)

func ExampleDownloadBill() {
	var ctx *wxpay.WXPay

	params := map[string]string{
		"appid": "APPID",
		// ...
	}
	resp, err := download.DownloadBill(ctx, params)

	fmt.Println(resp, err)
}
