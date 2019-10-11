/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package backend

import (
	"fmt"
	"time"

	"github.com/webx-top/echo"

	"github.com/admpub/nging/application/handler"
	"github.com/admpub/nging/application/middleware"
	"github.com/admpub/nging/application/registry/upload"
	"github.com/admpub/nging/application/registry/upload/table"
)

func init() {
	// 后台用户头像上传
	upload.CheckerRegister(`user-avatar`, func(ctx echo.Context, tis table.TableInfoStorer) (subdir string, name string, err error) {
		userID := ctx.Formx(`userId`).Uint64()
		if user := handler.User(ctx); user != nil {
			err = middleware.CheckAnyPerm(ctx, `manager/user_add`, `manager/user_edit`)
			if err != nil {
				return
			}
		} else {
			err = ctx.E(`请先登录`)
			return
		}
		timestamp := ctx.Formx(`time`).Int64()
		// 验证签名（避免上传接口被滥用）
		if ctx.Form(`token`) != upload.Token(`userId`, userID, `time`, timestamp) {
			err = ctx.E(`令牌错误`)
			return
		}
		if time.Now().Local().Unix()-timestamp > upload.UploadLinkLifeTime {
			err = ctx.E(`上传网址已过期`)
			return
		}
		if userID > 0 {
			name = `avatar`
		}
		subdir = fmt.Sprint(userID) + `/`
		tis.SetTableID(fmt.Sprint(userID))
		tis.SetTableName(`user`)
		tis.SetFieldName(`avatar`)
		return
	})
}
