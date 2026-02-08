class Solution(object):
    def isHappy(self, n, seen=None):
        """
        :type n: int
        :rtype: bool
        """

        if seen is None:
            seen = set()

        if n == 1:
            return True

        if n in seen:
            return False

        seen.add(n)

        total = 0
        while n > 0:
            digit = n % 10
            total += digit * digit
            n //= 10

        return self.isHappy(total, seen)
