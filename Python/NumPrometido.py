def numPrometido(num1,num2,num3,suma: int) -> bool: 
    if (num3 == 0): 
        if ((num2 + 1) == suma):
            return True
        else: 
            return False
    else: 
        if ((num1 % num3) == 0): 
            return numPrometido(num1,num2,num3-1,suma+num3)
        else: 
            return numPrometido(num1,num2,num3-1,suma)
        

for i in range(47,76):
    for j in range(i+1,77):
        if numPrometido(i,j,i-1,0) and numPrometido(j,i,j-1,0):
            print("El par de numeros es prometido: ", "(",i,",",j,")")
