class Solution:
    def xorOperation(self, n: int, start: int) -> int:
        l = []
        for i in range(n):
            l.append(start + i * 2)
        result = l[0]
        for i in range(1,n):
            result  = result ^ l[i]
        return result

