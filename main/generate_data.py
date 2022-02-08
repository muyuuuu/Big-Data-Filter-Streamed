# 编号，性别，考核得分
import random
from tqdm import tqdm

for i in tqdm(range(10000000)):
    string = []
    id_ = str(i).rjust(7, "0")
    gender = "1" if random.random() > 0.5 else "0"
    score = str(random.randint(80, 100))
    string.append(id_)
    string.append(gender)
    string.append(score)
    with open("data.txt", "a+") as f:
        f.write(",".join(string) + "\n")
