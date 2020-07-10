#https://leetcode-cn.com/problems/longest-palindrome/
class Solution:
    def longestPalindrome(self, s: str) -> int:
        dic = {}
        result = 0
        for ch in s:
            count = dic.get(ch,0)
            if count == 0:
                dic[ch] = 1
                continue
            if count == 1:
                dic[ch] = 0
                result = result + 1
        result = result * 2
        if result  < len(s):
            result = result  + 1
        return result 
