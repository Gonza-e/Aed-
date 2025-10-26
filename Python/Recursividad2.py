numeros = [42,53,65,23,12,54,23,28,63,62]

def maximo(a,max,indice) -> int:
    if indice == 10:
        return max 
    else: 
        if a[indice] > max:
            return maximo(a,a[indice],indice+1)
        else:
            return maximo(a,max,indice+1)


print("El maximo valor en el arreglo es:",maximo(numeros,numeros[0],0))