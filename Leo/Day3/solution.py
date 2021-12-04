import numpy as np

def initiate_list(l):
    nbr_ones = []
    for n in range(l):
        nbr_ones.append(0)
    return nbr_ones

def create_binaries(nbr_ones, lim):
    nbr1 = []
    nbr2 = []
    for n in nbr_ones:
        if n > lim:
            nbr1.append('1')
            nbr2.append('0')
        else:
            nbr1.append('0')
            nbr2.append('1')
    return "".join(nbr1), "".join(nbr2)

def main():
    file = open("data.txt")
    lines = file.readlines()

    nbr_ones = initiate_list(len(lines[0].strip()))

    nbr_lines = 0

    for line in lines:
        nbr = line.strip()
        for n,d in enumerate(nbr):

            if d == '1':
                nbr_ones[n] += 1
        nbr_lines += 1

    bin1,bin2 = create_binaries(nbr_ones, nbr_lines/2)

    print(bin1)

    dec1 = int(bin1,2)
    dec2 = int(bin2,2)

    sol = dec1*dec2
    print(sol)

#Recursive function for task2
#Majority is a boolean for whether looking for majority or not
def rec_nums(lines, ind, majority):
    if len(lines) <= 1:
        return lines
    ones = []
    nbr_ones = 0
    for line in lines:
        if line[ind] == '1':
            ones.append(True)
            nbr_ones += 1
        else:
            ones.append(False)

    if not majority:
        ones = np.logical_not(ones)

    if nbr_ones >= len(lines)/2:
        return rec_nums(lines[ones], ind+1, majority)
    else:
        return rec_nums(lines[np.logical_not(ones)], ind+1, majority)



def main2():
    lines = np.genfromtxt('data.txt',dtype='str')

    bin1 = rec_nums(lines,0,True)
    bin2 = rec_nums(lines,0,False)

    dec1 = int(bin1[0],2)
    dec2 = int(bin2[0],2)
    sol = dec1*dec2
    print(sol)

if __name__ == '__main__':
    main2()
