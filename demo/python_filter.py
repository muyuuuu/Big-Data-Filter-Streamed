import numpy as np
from mpi4py import MPI
import random
import time

comm = MPI.COMM_WORLD
rank = comm.Get_rank()
root = 0
nprocs = comm.Get_size()

if rank == 0:
    start = time.time()
    with open("data.txt", "r") as f:
        data = f.readlines()
        data = [float(i.strip()) for i in data]
        ave, res = divmod(len(data), nprocs)
        counts = [ave + 1 if p < res else ave for p in range(nprocs)]
        starts = [sum(counts[:p]) for p in range(nprocs)]
        ends = [sum(counts[: p + 1]) for p in range(nprocs)]
        data = [data[starts[p] : ends[p]] for p in range(nprocs)]
else:
    data = None


def is_large(n):
    return n >= 0.5


tmp = comm.scatter(data, root=0)
sendbuf = list(filter(is_large, tmp))
sendbuf = np.array(sendbuf)

sendcounts = np.array(comm.gather(len(sendbuf), root))

if rank == root:
    print("sendcounts: {}, total: {}".format(sendcounts, sum(sendcounts)))
    recvbuf = np.empty(sum(sendcounts), dtype=int)
else:
    recvbuf = None

comm.Gatherv(sendbuf=sendbuf, recvbuf=(recvbuf, sendcounts), root=root)
if rank == root:
    print("Gathered array: {}, {}s".format(recvbuf.size, time.time() - start))
