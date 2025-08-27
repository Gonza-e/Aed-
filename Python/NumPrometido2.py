def numPrometido(num1,num2,num3,num4,suma1,suma2) -> bool:
    if (num2 == 0) and (num4 == 0):
        if (suma1 == num3 + 1) and (suma2 == num1 + 1):
            return True
        else:
            return False
    else: 
        if num2 > 0:
            if (num1 % num2) == 0:
                return numPrometido(num1,num2-1,num3,num4,suma1+num2,suma2)
            else:
                return numPrometido(num1,num2-1,num3,num4,suma1,suma2)
        
        if num4 > 0:
            if (num3 % num4) == 0:
                return numPrometido(num1,num2,num3,num4-1,suma1,suma2+num4)
            else:
                return numPrometido(num1,num2,num3,num4-1,suma1,suma2)
            
for i in range(47,74):
    for j in range(i+1,76):
        if numPrometido(i,i-1,j,j-1,0,0):
            print("El par de numeros es prometido: ","(",i,",",j,")")

        