class Solution:
    def countGoodTriplets(self, arr: List[int], a: int, b: int, c: int) -> int:
        length = len(arr)
        count = 0
        i = 0
        while i < length:
            j = i + 1
            while j < length:
                k = j + 1
                while k < length:
                    if abs(arr[i]-arr[j]) <= a and abs(arr[j]-arr[k]) <= b and abs(arr[i]-arr[k]) <= c:
                        count = count + 1
                    k = k + 1
                j = j + 1
            i = i + 1
        return count
