class Solution:
    def largestOddNumber(self, num: str) -> str:
        s = str(abs(int(num)))
        i = len(s) - 1
        result = ""
        while i >= 0:
            n = int(s[i])
            if not len(result):
                if n % 2 == 1:
                    result = str(n)
            else:
                result = str(n) + result
            i = i - 1
        return result


s = Solution()
s.largestOddNumber("52")
