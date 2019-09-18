import matplotlib.pyplot as plt 
import ctypes


def load(x : list, y : list, c : list = []) -> dict:
    args = {
        "x": x,
        "y": y,
        "c": c
        }
    return args


def plot(args : dict, title : str, fn):
    n = len(args["x"])
    y = [ "{0} ({1}, {2})".format(title, args["x"][i], args["y"][i]) for i in range(n)]
    x = [fn(ctypes.c_int(args["x"][i]), ctypes.c_int(args["y"][i])) for i in range(n)]
    plt.plot(x, y, color="c", linestyle="solid", marker="o", markersize=5)
    if args["c"] != []:
         plt.plot(args["c"], y, color="r", linestyle="dashed", marker="o", markersize=9)
    plt.title(title)
    plt.grid()
    plt.show()

