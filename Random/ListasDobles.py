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

        if self.prim is None:   #Lista vacia (si prim = nil entonces)
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
                if p is None:   #Insertar al final (si p = nil entonces)
                    self.ult.prox = nuevo 
                    nuevo.prox = None
                    nuevo.ant = self.ult 
                    self.ult = nuevo 
                else:   #Insertar en el medio 
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
            print(f"{p.num}",end = " <-> ")
            p = p.prox
        print("nil \n")
    
            
#De aca en adelante son solamente los bucles para poder probar las cargas y eliminaciones

lista = ListaDoble()
bandera = True
while bandera:
    numero = int(input("Ingrese un numero mayor a cero \n"))
    if numero >= 0:
        lista.cargaOrdenada(numero)
    else: 
        bandera = False

lista.mostrarNodo()

bandera = True
while bandera:
    numero1 = int(input("Ingrese valor que desea eliminar: \n"))
    if numero1 >= 0:
        lista.eliminarNodo(numero1)
    else:
        bandera = False

    if lista.prim == None:
        bandera = False
    lista.mostrarNodo()
