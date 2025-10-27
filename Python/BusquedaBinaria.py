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
        
numeros = [2,5,6,14,45,67,70]
medio = (0 + 6) // 2

valor = int(input("Ingrese un valor para encontrar \n"))

print("Valor encontrado:",busquedaBinaria(numeros,0,6,medio,valor))