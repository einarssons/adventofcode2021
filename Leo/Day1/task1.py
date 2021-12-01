import numpy as np

#For part2
def cumsum_depths(w,depths):
    sums = np.cumsum(depths)
    sums[w:] = sums[w:]-sums[:-w]
    return sums[w-1:]

def main():
    depths = np.loadtxt('data.txt')
    #depths = np.array([199, 200, 208, 210, 200, 207, 240, 269, 260, 263])
    count = 0
    sum_depths = cumsum_depths(3,depths)
    for n,d in enumerate(sum_depths):
        if n is not 0:
            if sum_depths[n-1] < d:
                count += 1
    print(count)

if __name__ == "__main__":
    main()
