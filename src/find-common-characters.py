#https://leetcode-cn.com/problems/find-common-characters/
from typing import List
class Solution:
    def mix(self, str1 :str,str2: str) ->str:
        if len(str1) == 0 || len(str2) == 0
            return ""
        dic = {}
        result = ""
        for ch in str1:
            count = dic.get(ch,0) + 1
            dic[ch]= count
        for ch in str2:
            count = dic.get(ch,0)
            if count > 0 :
                result = result + ch
                count = count -1
                dic[ch]=count
        return result

    def commonChars(self, A: List[str]) -> List[str]:
        if len(A) == 0:
            return []
        result = A[0]
        for s in A:
            result = self.mix(result,s)
            print(result)
        return [ch for ch in result]
s = Solution()
s.commonChars(["acabcddd","bcbdbcbd","baddbadb","cbdddcac","aacbcccd","ccccddda","cababaab","addcaccd"])
