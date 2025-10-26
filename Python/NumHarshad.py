def numHarshad(num,num1,suma: int) -> bool:
    if num1 == 0:
        if (num % suma) == 0:
            return True 
        else:
            return False 
    else: 
        return numHarshad(num,num1 // 10,suma+(num1 % 10))
        

n = 18 

if numHarshad(n,n,0):
    print("El numero es Harshad")
else: 
    print("El numero no es Harshad")