/*
 *   Copyright silenceper/wechat Author(https://silenceper.com/wechat/). All Rights Reserved.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 *
 *   You can obtain one at https://github.com/silenceper/wechat.
 *
 */

// Package minidrama Mini Program entertainment mini-drama related interface
package minidrama

import (
	"github.com/silenceper/wechat/v2/miniprogram/context"
)

// NewMiniDrama 实例化小程序娱乐直播 API
func NewMiniDrama(ctx *context.Context) *MiniDrama {
	return &MiniDrama{
		ctx: ctx,
	}
}
