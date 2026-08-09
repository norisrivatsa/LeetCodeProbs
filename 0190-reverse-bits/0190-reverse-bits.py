class Solution(object):
    def reverseBits(self, n):
        """
        :type n: int
        :rtype: int
        """
        bits = []

        for i in range(32):
            bits.append(n % 2)
            n //= 2

        number = 0

        for bit in bits:
            number = number * 2 + bit

        return number
