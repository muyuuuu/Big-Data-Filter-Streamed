import time

start = time.time()
with open("data.txt", "r") as f:
    data = f.readlines()
    for i in data:
        item = float(i.strip())
        if item >= 0.5:
            with open("result.txt", "a+") as f:
                f.writelines(i)
print(time.time() - start)
