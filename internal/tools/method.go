/*
 * @Date: 2026-05-31 15:00:14
 * @LastEditTime: 2026-05-31 15:50:58
 * @FilePath: /dark_back/internal/tools/method.go
 * @Description:
 */
package tools

import (
	"net/netip"
	"regexp"
	"strings"
)

/**
 * @description:转小写 并且只保留字母与数字
 * @param {string} s
 * @return {*}
 */
func ToLowerAndRemoveSpecialChars(s string) string {
	// 1. 转小写
	s = strings.ToLower(s)

	// 2. 正则：只保留 字母 和 数字，其他全部删掉
	reg := regexp.MustCompile(`[^a-z0-9]`)
	return reg.ReplaceAllString(s, "")
}

/**
 * @description: 格式化IP
 * @param {string} ipStr
 * @return {*}
 */
func FormatIP(ipStr string) string {
	// 1. 解析 IP（自动支持 ipv4/ipv6）
	addr, err := netip.ParseAddr(ipStr)
	if err != nil {
		// 解析失败，直接清理非法字符
		return cleanIPString(ipStr)
	}

	// 2. IPv6 格式化（去掉端口、zone、%）
	if addr.Is6() {
		// 返回不带 : 的纯字符串
		return strings.ReplaceAll(addr.String(), ":", "")
	}

	// 3. IPv4 格式化（去掉 .）
	return strings.ReplaceAll(addr.String(), ".", "")
}

/**
 * @description: 清理非法字符
 * @param {string} s
 * @return {*}
 */
func cleanIPString(s string) string {
	var builder strings.Builder
	for _, c := range s {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
