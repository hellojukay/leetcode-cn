class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        result = numBottles
        bottles = numBottles
        while bottles >= numExchange:
            result = result + int(bottles/numExchange)
            bottles = int(bottles/numExchange) + (bottles % numExchange)
        return result

