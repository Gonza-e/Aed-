def busquedaBinaria(arr,izq,der,medio,num) -> int:
    if izq == der:
        return 0
    else:
        if num == arr[medio]:     
            return arr[medio] 
        else:
            if num > arr[medio]:
                return busquedaBinaria(arr,medio,der,(izq+der)//2,num)
            else:
                return busquedaBinaria(arr,izq,medio,(izq+der)//2,num)
        
numeros = [30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]
medio = (numeros[0] + numeros[29]) // 2

print(busquedaBinaria(numeros,numeros[0],numeros[29],medio,2))