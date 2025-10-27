A = [[23, 44, 37, 45, 62, 23, 23],[28, 62, 66, 16, 72, 36, 25],[23,2,24,21,12,23,12]]


for i in range(len(A)-1):
    for j in range(len(A[i])):
        print("Piloto:", i , "Equipo:", j)
        print("Victorias:",A[2][j])

        dif = A[0][j] - A[1][j]
        if dif < 0:
            dif = dif * (-1)
        
        print("Diferencia de puntos entre ambos corredores:",dif,"\n")
        dif = 0


