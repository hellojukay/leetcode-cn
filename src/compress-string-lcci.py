class Solution:
    def compressString(self, S: str) -> str:
        if not S:
            return S
        pre = ''
        count = 0
        result = ""
        for ch in S:
            if pre == ch:
                count = count + 1
            else:
                if count > 0:
                    result = result + pre + str(count)
                count = 1
            pre = ch
        result = result + pre + str(count)
        if len(result) >= len(S):
            return S
        return result
