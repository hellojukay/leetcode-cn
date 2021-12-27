class Solution:
    def kthDistinct(self, arr: List[str], k: int) -> str:
        m = {}
        for ch in arr:
            n = m.get(ch, 0)
            m[ch] = n + 1
        n = 0
        for ch in arr:
            if m.get(ch, 0) == 1:
                n = n + 1
            if n == k:
                return ch
        return ""
