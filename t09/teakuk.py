# %%
# 태국의 
import numpy as np
import matplotlib.pyplot as plt

REPEAT = 10000
TOTAL_PERSON = 100
HELL_CHOSEN = 20
HELL_PERCENT = 0.3

def limit(line):
    count = 0
    for i, v in enumerate(line):
        if v:
            count += 1
        if count > HELL_CHOSEN:
            break
    newLine = np.copy(line)
    newLine[i:] = False
    return newLine

            

result = np.random.choice(a=[True, False], p=[HELL_PERCENT, 1- HELL_PERCENT], size=TOTAL_PERSON * REPEAT).reshape((REPEAT, TOTAL_PERSON))
vecRes = np.vectorize(limit, signature='(n)->(m)')(result)
# n 번째 인원
x = np.arange(TOTAL_PERSON) 
# n 번째 인원이 당첨될 확률
y = np.vectorize(lambda i: sum(vecRes[:, i]) / REPEAT, otypes=[np.float64])(x)

plt.plot(x, y)
plt.show()