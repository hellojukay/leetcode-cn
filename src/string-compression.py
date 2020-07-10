# https://leetcode-cn.com/problems/string-compression/
from typing import List;
class Solution:
    def compress(self, chars: List[str]) -> int:
        if len(chars) <= 1:
            return len(chars)
        pre = chars[0]
        count = 1
        index = 0
        for i in range(1,len(chars)):
            if pre == chars[i]:
                count = count+1
            else:
                if count == 1:
                    chars[index] = pre
                    index = index + 1
                else:
                    chars[index] = pre
                    index = index + 1
                    for ch in  str(count):
                        chars[index] = ch
                        index = index + 1
                count = 1
                pre = chars[i]
        if count == 1:
            chars[index] = pre
            return index + 1 
        else:
            chars[index] = pre
            index = index + 1
            for ch in  str(count):
                chars[index] = ch
                index = index + 1
            return index

