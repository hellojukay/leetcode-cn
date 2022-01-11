import re
class Solution:
    def complexNumberMultiply(self, num1: str, num2: str) -> str:
        a,b,_ = re.split("[+i]",num1)
        c,d,_ = re.split("[+i]",num2)
        s = (int(a) * int(c))- (int(b) * int(d))
        x = (int(a) * int(d)) + (int(b) * int(c))
        return  str(s) + "+" + str(x) + "i"


