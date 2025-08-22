def calcularPromedio(arr,i) -> float:
    if i == 9: 
        return arr[i] / 10 
    else: 
        return arr[i] / 10 + calcularPromedio(arr,i+1)
    
numeros = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]  # 10 lugares
print("El promedio es:", calcularPromedio(numeros,0))