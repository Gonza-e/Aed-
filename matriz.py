a = [
    [1,5,3,4,13],
    [8,6,1,4,19],
    [3,8,9,3,23],
    [12,19,13,11,110]
]

def obtenerFranja(x) -> str:
    if x == 0: return "8-10"
    elif x == 1: return "10-12"
    elif x == 2: return "12-14"
    elif x == 3: return "14-16"

for i in range(len(a)):
    for j in range(len(a[i])):

        if (i != 3) and (j != 4):
            print("Caja",i,"Franja:",obtenerFranja(j),"Monto:",a[i][j],"$")
        else: 
            if (i == 3) and (j != 4):
                print("Total por caja:",a[i][j],"$")
            else:
                if (j == 4) and (i != 3):
                    print("Total por franja:",a[i][j],"$")
                else:
                    if (j == 4) and (i == 3):
                        print("Total general:",a[i][j],"$")
