class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) == 0:
            return ""
        index = 0
        length = len(s)
        m = 0
        result = s[0]
        while index < length:
            start, end = index, index
            # 基数对称
            while start >= 0 and end < length and s[start] == s[end]:
                start = start - 1
                end = end + 1
            if end - start > m:
                m = end - start
                result = s[start+1:end]
            # 偶数对称情况
            start, end = index, index+1
            while start >= 0 and end < length and s[start] == s[end]:
                start = start - 1
                end = end + 1
            if end - start > m:
                m = end - start
                result = s[start+1:end]
            index += 1
        return result


s = Solution()
print(s.longestPalindrome('a'))
print(s.longestPalindrome('bb'))
print(s.longestPalindrome('bbb'))
print(s.longestPalindrome('abcddcb'))
print(s.longestPalindrome('abcdadcb'))
