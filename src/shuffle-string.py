class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        m = {}
        length = len(indices)
        for n in range(0, length):
            x = indices[n]
            m[x] = s[n]
        result = ""
        for i in range(0, len(s)):
            result = result + m[i]
        return result
