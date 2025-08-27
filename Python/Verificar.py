def verificar(arr,i) -> str: 
    if i == 30:
       return "Fin del proceso"
    else: 
        if i == arr[i]:
            print("Coinciden: ","(",arr[i],",",i,")")
        else:
            print("No coinciden")
        
        return verificar(arr,i+1)
    
numeros = [30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1]

print(verificar(numeros,0))

"""Dado un vector de números enteros ordenados decrecientemente, diseñar una función recursiva que devuelva el
resultado de comprobar si el valor de alguno de los elementos del vector coincide con su indice
"""
