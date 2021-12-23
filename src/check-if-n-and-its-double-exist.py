class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        s = set()
        for n in arr:
            if n*2 in s:
                return True
            else:
                if n % 2 == 0:
                    if int(n/2) in s:
                        return True
                s.add(n)
        return False
