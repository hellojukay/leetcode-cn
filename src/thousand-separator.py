class Solution:
    def thousandSeparator(self, n: int) -> str:
        return  str(re.sub(r'(\d{3})',r'\g<1>.',str(n)[::-1])[::-1]).lstrip('.')
