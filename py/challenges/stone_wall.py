import unittest

# You are going to build a stone wall. The wall should be straight and N meters long, 
# and its thickness should be constant; however, it should have different heights 
# in different places. The height of the wall is specified by an array H of N positive 
# integers. H[I] is the height of the wall from I to I+1 meters to the right of its
# left end. In particular, H[0] is the height of the wall's left end and H[N−1] is 
# the height of the wall's right end.

# The wall should be built of cuboid stone blocks (that is, all sides of such blocks 
# are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

# Write a function:

# def solution(H)

# that, given an array H of N positive integers specifying the height of the wall, 
# returns the minimum number of blocks needed to build it.

# For example, given array H containing N = 9 integers:

#   H[0] = 8    H[1] = 8    H[2] = 5
#   H[3] = 7    H[4] = 9    H[5] = 8
#   H[6] = 7    H[7] = 4    H[8] = 8
# the function should return 7. The figure shows one possible arrangement of seven blocks.



# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [1..100,000];
# each element of array H is an integer within the range [1..1,000,000,000].


N_RANGE = [1,100000]
H_RANGE = [1,1000000000]

def solution(A):
    intArr = A

    count = 1
    for i in range (1, len(intArr)):
        prevHeight = intArr[i-1]
        newHeight = intArr[i]
        if newHeight > prevHeight:
            count += 1
        elif newHeight < prevHeight:
            counted = False
            for j in range (0, i): # check if the new height is already counted in a previous block
                jHeight = intArr[j]
                if newHeight > jHeight:                            
                    break
                elif newHeight == jHeight:
                    counted = True
                    break
            if not counted:
                count += 1
    return count

    
class StoneWallTest(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution([8,8,5,7,9,8,7,4,8]), 7)
    
    def test_min(self):
        self.assertEqual(solution([1]), 1)
        self.assertEqual(solution([1,1]), 1)
        self.assertEqual(solution([1,1,1]), 1)

    def test_max(self):
        intArr = [n for n in range(*N_RANGE)]
        self.assertEqual(solution(intArr), N_RANGE[1]-1)
        intArrStep = [n*1000 for n in range(*N_RANGE)]
        self.assertEqual(solution(intArrStep), N_RANGE[1]-1)


