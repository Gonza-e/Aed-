def verificar(A,ind) -> int:
    if ind == 5:
        return 0
    else:
        if A[ind] >= 6:
            return 1 + verificar(A,ind+1)
        else:
            return verificar(A,ind+1)
        

Arreglo = [4,6,7,3,1]

if verificar(Arreglo,0) >= 3:
    print("Aprobado")
else: 
    print("Desaprobado")

