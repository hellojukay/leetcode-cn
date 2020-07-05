class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        max = 0
        result = []
        for n in candies:
            if n >= max :
                max = n

        for index in range(len(candies)):
            if candies[index] + extraCandies >= max:
                result.append(True)
            else:
                result.append(False)
        return result
