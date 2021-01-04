//  Copyright 2020 Go-Ceres
//  Author https://github.com/go-ceres/go-ceres
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package CeresError

import "encoding/json"

// Error 实现error包接口
func (e *Error) Error() string {
	d, _ := json.Marshal(e)
	return string(d)
}

func (e *Error) WithId(id string) *Error {
	e.Id = id
	return e
}

// WithAid 设置应用id
func (e *Error) WithAid(aid string) *Error {
	e.Aid = aid
	return e
}

// WithTid 设置请求trace id
func (e *Error) WithTid(tid string) *Error {
	e.Tid = tid
	return e
}

// WithPkg 设置出错的包
func (e *Error) WithPkg(pkg string) *Error {
	e.Pkg = pkg
	return e
}

// New创建一个基本的错误信息
func New(msg string, code ...int32) *Error {
	e := &Error{
		Msg: msg,
	}
	if len(code) > 0 {
		e.Code = code[0]
	}
	return e
}

// Parse 解析成结构体
func Parse(err string) *Error {
	e := new(Error)
	res := json.Unmarshal([]byte(err), e)
	if res != nil {
		e.Msg = err
	}
	return nil
}
