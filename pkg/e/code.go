//Package e ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 14:53:04
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:41:23
 */
package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_NICK           = 10001
	ERROR_EXIST_USER           = 10002
	ERROR_NOT_EXIST_USER       = 10003
	ERROR_NOT_COMPARE          = 10004
	ERROR_NOT_COMPARE_PASSWORD = 10005
	ERROR_FAIL_ENCRYPTION      = 10006
	ERROR_NOT_EXIST_PRODUCT    = 10007
	ERROR_NOT_EXIST_ADDRESS    = 10008
	ERROR_EXIST_FAVORITE       = 10009

	ERROR_AUTH_CHECK_TOKEN_FAIL       = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT    = 20002
	ERROR_AUTH_TOKEN                  = 20003
	ERROR_AUTH                        = 20004
	ERROR_AUTH_INSUFFICIENT_AUTHORITY = 20005

	ERROR_DATABASE = 30001

	ERROR_OSS = 40001
)
