class Solution:
    def canMakeArithmeticProgression(self, arr: List[int]) -> bool:
        arr.sort()
        if len(arr) <= 2:
            return True
        dif = arr[1] - arr[0]
        for n in range(len(arr)-1):
            if arr[n+1] != arr[n] + dif:
                return False
        return True

