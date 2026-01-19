def invertir(A,B,indA,indB):
    if indA < 10:
        invertir(A,B,indA + 1,indB - 1)

    B[indB] = A[indA]


Arr1 = [1,2,3,4,5,6,7,8,9,10]
Arr2 = [0,0,0,0,0,0,0,0,0,0]

invertir(Arr1,Arr2,0,9)

for i in range(0,10):
    print("Arreglo 1:",Arr1[i],"Arreglo 2:",Arr2[i])

