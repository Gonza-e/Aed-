def indice(A:list,let:str) -> int:
    for i in range(5):
        if A[i] == let:
            return i 


Arreglo = ["A","E","I","O","U"]

reemplazo = (indice(Arreglo,"U")-2) % 5 

print(Arreglo[reemplazo])

