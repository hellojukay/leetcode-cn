class Solution:
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        def check(n: int) -> bool:
            s = str(n)
            if "i" in s:
                return False
            for ch in s:
                if n % int(ch) != 0:
                    return False
            return True
        result = []
        for n in range(left, right+1):
            if check(n):
                result.append(n)
        return result
