#https://leetcode-cn.com/problems/first-unique-character-in-a-string/
class Solution:
    def firstUniqChar(self, s: str) -> int:
        dic = {}
        for ch in s:
            count = dic.get(ch,0)
            dic[ch] = count + 1
        for i,ch in enumerate(s):
            count = dic.get(ch)
            if count == 1:
                return i
        return -1

