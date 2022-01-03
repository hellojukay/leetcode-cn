class Solution:
    def checkString(self, s: str) -> bool:
        if len(s) <= 1:
            return True
        flag = 1
        pre = s[0]
        for ch in s:
            if pre != ch:
                flag = 0
            if not flag and ch == pre:
                return False
            if pre == "b" and ch == "a":
                return False
        return True
