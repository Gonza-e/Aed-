def verificar(arr,i,suma) -> bool:
    if i == -1: 
        if (suma % 10) == 0:
            return True
        else:
            return False
    else: 
        if (i % 2) != 0: 
            arr[i] = arr[i] * 2
            if arr[i] > 9:
                arr[i] = (arr[i] % 10) + (arr[i] // 10)
            return verificar(arr,i-1,suma + arr[i])
        else:
            return verificar(arr,i-1,suma)

tarjeta = [4,7,3,9,6,1,7,4,7,3,7,8,6,7,1,4] 

if verificar(tarjeta,15,0):
    print("La tarjeta es valida")
else: 
    print("La tarjeta no es valida")
    
#Ejercicio de recursividad sacado del 3er parcial de 2024 (Parcial3/Parcial2024B.go)