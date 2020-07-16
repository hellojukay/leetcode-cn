# https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/
class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        flag = False
        queue = []
        pre = ''
        count = 0
        for ch in s:
            if pre == ch:
                count = count + 1
            else:
                if count > 0:
                    queue.append((pre,count))
                count = 1
            pre = ch

        queue.append((pre,count))
        s = ""
        for (ch,count) in queue:
            n = count
            if count >= k:
                n = count -k
                flag = True
            for i in range(n):
                s = s + ch
        if not flag:
            return s
        return self.removeDuplicates(s,k)

