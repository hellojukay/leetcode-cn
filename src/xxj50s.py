class Solution:
    def reverseString(self, s: List[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        if len(s) < 2:
            return s
        start, end = 0, len(s)-1
        while True:
            if start >= end:
                break
            tmp = s[start]
            s[start] = s[end]
            s[end] = tmp
            start = start + 1
            end = end - 1
        return None
