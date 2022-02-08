import random
from tqdm import tqdm

data = []
for i in tqdm(range(10000)):
    with open("data.txt", 'a+') as f:
        f.write(str(random.random()) + '\n')
