/*
 * @Author: lihao lihao@ikbvip.com
 * @Date: 2023-05-04 14:38:44
 * @LastEditors: lihao lihao@ikbvip.com
 * @LastEditTime: 2024-04-28 15:12:10
 * @FilePath: \iec-em-cem-server\common\rediskey\rediskey.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package rediskey

const DefaultDelimiter = ":"

//// 所有本服务redisKey定义地方 规则服务名字:xx:xx:xx
//const ServiceCemCacheKey = "__IEC:CEM" //缓存头
//
//// redis缓存雪花算法机器id
//var SnowflakeWorkIdKey = func(names ...string) string {
//	return redisUtil.CreateKey(false, "")(ServiceCemCacheKey, "snowflake", "workId", strings.Join(names, DefaultDelimiter))
//}
