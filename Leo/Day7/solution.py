import numpy as np

def setup_zero(crabs):
    #zero_pos = np.zeros(np.max(crabs), dtype=int)
    zero_pos = np.array([],dtype='int')
    for c in crabs:
        zero_pos = np.append(zero_pos,0-c)
    return zero_pos

def task1():
    crabs = np.genfromtxt('small_data.txt',dtype='int', delimiter = ',')

    crab_pos = setup_zero(crabs)
    min_fuel = np.abs(crab_pos).sum()

    for n in range(np.max(crabs)):
        crab_pos = crab_pos+1
        min_fuel = min(np.abs(crab_pos).sum(),min_fuel)

    print(min_fuel)

#For task2 - create list of triagle numbers ;) 
def create_cumsum(size):
    cumsums = np.array([0])
    for m in range(1,size+1):
        cumsums = np.append(cumsums,cumsums[m-1]+m)
    return cumsums

def task2():
    crabs = np.genfromtxt('data.txt',dtype='int', delimiter = ',')
    cumsums = create_cumsum(max(crabs))
    crab_pos = setup_zero(crabs)
    min_fuel = (cumsums[np.abs(crab_pos)]).sum()

    for n in range(np.max(crabs)):
        crab_pos = crab_pos+1
        min_fuel = min((cumsums[np.abs(crab_pos)]).sum(),min_fuel)

    print(min_fuel)


if __name__ == "__main__":
    task2()
