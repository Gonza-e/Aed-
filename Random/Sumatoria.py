def Sumatoria(n,i,suma: int) -> int:
    if n == 0:
        return suma
    else:
        return Sumatoria(n - 1,i + 1,suma + (3 * i + 1))
    

print(Sumatoria(3,1,0))