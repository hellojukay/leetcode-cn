class Solution:
    def addDigits(self, num: int) -> int:
        if num < 10:
            return num
        tmp = 0
        while num != 0 :
            tmp = tmp + num % 10
            num = int(num /10)
        return self.addDigits(tmp)

