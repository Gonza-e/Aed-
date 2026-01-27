class Nodo:
    def __init__(self,letra):
        self.letra = letra 
        self.prox = None 
    
class Lista: 
    def __init__(self):
        self.prim = None 
        self.ult = None 
    
    def cargaEncolada(self,letra):
        q = Nodo(letra)
        p = self.prim

        if self.prim is None:
            q.prox = self.prim
            self.prim = q 
            self.ult = q 
        else:
            q.prox = self.ult.prox 
            self.ult.prox = q 
            self.ult = q 
    
    def mostrarLista(self):
        p = self.prim
        fras = ""
        while p is not None:
            fras += p.letra
            p = p.prox  

        print(fras)

    def encriptar(self):
        p = self.prim
        letra = 0
        while p is not None:
            if p.letra != ' ':
                letra = AR.index(p.letra)
                letra = (letra + 1) % 26 
                p.letra = AR[letra]
            p = p.prox

    def desencriptar(self):
        p = self.prim
        letra = 0
        while p is not None:
            if p.letra != ' ':
                letra = AR.index(p.letra)
                letra = (letra - 1) % 26 
                p.letra = AR[letra]
            p = p.prox

AR = ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]

frase = "ATAQUE AL AMANECER"
lista = Lista()

for letra in frase:
    lista.cargaEncolada(letra)

lista.mostrarLista()
lista.encriptar()
lista.mostrarLista()

while True: 
    respuesta = input("Ingrese si o no: \n")

    if respuesta == "si":
        lista.desencriptar()
        lista.mostrarLista()

    break




