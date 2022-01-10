class Solution:
    def makeFancyString(self, s: str) -> str:
        return re.sub(r'(.)\1{2,}',r'\g<1>\g<1>',s)

