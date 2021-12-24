class Solution:
    def countSubstrings(self, s: str) -> int:
        result = 0
        i = 0
        length = len(s)
        while i < length:
            start, end = i, i
            while start >= 0 and end < length and s[start] == s[end]:
                result = result + 1
                start = start - 1
                end = end + 1
            i = i + 1
        i = 0
        while i < length:
            start, end = i, i + 1
            while start >= 0 and end < length and s[start] == s[end]:
                result = result + 1
                start = start - 1
                end = end + 1
            i = i + 1
        return result


s = Solution()
print(s.countSubstrings("abc"))
