import numpy as np

#355386 Slow, slow, slooooooow
def task1():
    fish = np.genfromtxt('data.txt',dtype='int', delimiter = ',')
    days = 0
    while days < 80:
        fishies = len(fish)
        for n in range(fishies):
            if fish[n] == 0:
                fish[n] = 6
                fish = np.append(fish,8)
            else:
                fish[n] -= 1
        days += 1

    print(len(fish))


#Time to speed it up! And speedy it is
def task2():
    fish_init = np.genfromtxt('data.txt',dtype='int', delimiter = ',')
    fish = np.array([0]*9)
    for f in fish_init:
        fish[f] += 1

    days = 0
    while days < 256:
        birthFish = fish[0]
        for n in range(8):
            fish[n] = fish[n+1]
        fish[8] = birthFish
        fish[6] = fish[6] + birthFish
        days += 1

    print(np.sum(fish))

if __name__ == "__main__":
    task2()
