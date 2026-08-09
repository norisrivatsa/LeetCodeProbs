class Solution(object):
    def countBits(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        bit_array = []
        limit = n + 1
        i = 0
        def binary(i, bit_array):
            if i >= limit:
                return bit_array

            result = []

            temp = i

            while temp > 0:
                result.append(temp % 2)
                temp //= 2

            count = 0

            for j in result:
                if j == 1:
                    count += 1

            bit_array.append(count)

            return binary(i + 1, bit_array)

        return binary(i, bit_array)