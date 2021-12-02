import pandas as pd
import numpy as np

#Task 1
def main():
    df = pd.read_csv('smallData.csv', sep = ' ', header = None)
    front = df.iloc[df[df.iloc[:,0] == 'forward'].index,1].sum()
    up = df.iloc[df[df.iloc[:,0] == 'up'].index,1].sum()
    down = df.iloc[df[df.iloc[:,0] == 'down'].index,1].sum()
    sol = front*(down-up)
    print(sol)

#Task 2
def main2():
    df = pd.read_csv('data.csv', sep = ' ', header = None)
    forw  = 0
    depth = 0
    aim = 0

    for index, row in df.iterrows():
        com = row.iloc[0]
        if com == 'forward':
            forw += row.iloc[1]
            depth +=  row.iloc[1]*aim
        elif  com == 'down':
            aim += row.iloc[1]
        elif com == 'up':
            aim -= row.iloc[1]

    sol = forw*(depth)
    print(sol)


if __name__ == "__main__":
    main2()
