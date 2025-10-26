def numNarcisista(num1,num2,num3,num_og,cont1,cont2: int) -> bool:
    if num1 == 0:
        if cont1 == 0:
            if num3 == num_og:
                return True 
            else: 
                return False 
        else: 
            return numNarcisista(num1,num2 // 10,num3 + ((num2 % 10) ** cont2),num_og,cont1-1,cont2)
    else: 
        return numNarcisista(num1 // 10,(num2 * 10) + (num1 % 10),num3,num_og,cont1+1,cont2+1)
    

num = 9474

if numNarcisista(num,0,0,num,0,0):
    print("El numero es narcisista")
else:
    print("El numero no es narcisista")
