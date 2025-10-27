class Nodo:
    def __init__(self,numero):
        self.num = numero
        self.prox = None
        self.ant = None
    
class ListaDoble: 
    def __init__(self):
        self.prim = None
        self.ult = None

    def cargaOrdenada(self,numero):
        nuevo = Nodo(numero)

        if self.prim is None:
            self.prim = nuevo 
            self.ult = nuevo 
            nuevo.ant = None
            nuevo.prox = None
        else:
            p = self.prim
            while (p is not None) and (nuevo.num < p.num):
                p = p.prox 
            
            if self.prim == p:
                nuevo.ant = None
                nuevo.prox = self.prim 
                self.prim.ant = nuevo 
                self.prim = nuevo 
            else:
                if p is None:
                    self.ult.prox = nuevo 
                    nuevo.prox = None
                    nuevo.ant = self.ult 
                    self.ult = nuevo 
                else: 
                    (p.ant).prox = nuevo 
                    nuevo.ant = p.ant
                    nuevo.prox = p 
                    p.ant = nuevo

    def eliminarNodo(self,valor):
        if self.prim is None:
            print("Error lista vacia")
        else:
            p = self.prim 
            while (p is not None) and (valor != p.num):
                p = p.prox 

            if self.prim == self.ult:
                self.prim = None
                self.ult = None
            else: 
                if p == self.prim:
                    self.prim = self.prim.prox 
                    self.prim.ant = None
                else:
                    if p == self.ult:
                        self.ult = self.ult.ant 
                        self.ult.prox = None
                    else:
                        (p.ant).prox = p.prox 
                        (p.prox).ant = p.ant 
        e = p 
        del e 

    def mostrarNodo(self):
        p = self.prim
        while p is not None:        #Equivalente a poner Mientras p <> nil 
            print(f"Numero {p.num}","<->")
            p = p.prox
            

numero = 0
lista = ListaDoble()
while numero >= 0:
    numero = int(input("Ingrese un numero mayor a cero \n"))
    if numero > 0:
        lista.cargaOrdenada(numero)

lista.mostrarNodo()

numero1 = 0
while numero1 >= 0:
    numero1 = int(input("Ingrese valor que desea eliminar: \n"))
    if numero1 > 0:
        lista.eliminarNodo(numero1)
    lista.mostrarNodo()
