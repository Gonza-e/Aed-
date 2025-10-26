#La suma de sus dígitos es igual al producto de sus dígitos.
#Ejemplo: 1124 → (1+1+2+4 = 8) y (1×1×2×4 = 8) ✅

def numSpy(num,num1,sum,prod: int) -> bool:
    if num1 == 0:
        if sum == prod:
            return True
        else: 
            return False 
    else: 
        return numSpy(num,num1 // 10,sum + (num1 % 10),prod *(num1 % 10))
    

n = 1124 

if numSpy(n,n,0,1):
    print("El numero es espia")
else: 
    print("El numero no es espia")
