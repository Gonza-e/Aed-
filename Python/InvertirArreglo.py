def invertir(A:list,indA,x:int):
    if indA < 11:
        x = A[11-indA]
        invertir(A,indA+1,x)
        A[indA] = x 

Arr1 = [1,2,3,4,5,6,7,8,9,10]
Arr2 = {1:1,2:2,3:3,4:4,5:5,6:6,7:7,8:8,9:9,10:10}

#print(Arr1)
#invertir(Arr1,0,0)
#print(Arr1)

invertir(Arr2,1,0)
print(Arr2)


