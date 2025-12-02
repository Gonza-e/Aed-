class Nodo: 
    def __init__(self,dato):
        self.dato = dato
        self.prox = None 
        self.ant = None 
    
class ListaDoble: 
    def __init__(self):
        self.prim = None
        self.ult = None 
    
    def cargaEncolada(self,dato):
        q = Nodo(dato)
        if self.prim is None:
            self.prim = q 
            self.ult = q 
            q.prox = None
            q.ant = None 
        else: 
            q.prox = self.ult.prox 
            self.ult.prox = q 
            self.ult = q 
    
    def mostrarLista(self):
        p = self.prim 
        while (p is not None):
            print(f"{p.dato}", end=" <-> ")
            p = p.prox 
        print("nil")
    
    def esPalindromo(self,proximo,ultimo):
        if (proximo == ultimo) or (proximo.ant == ultimo):
            return True
        else:
            if proximo.dato == ultimo.dato:
                return self.esPalindromo(proximo.prox,ultimo.ant)
            else: 
                return False 
            

lista = ListaDoble()
palabra = input("Ingrese una palabra o numero para analizar \n")

for letra in palabra:
    lista.cargaEncolada(letra)

if lista.esPalindromo(lista.prim,lista.ult):
    print("Es palindromo")
else: 
    print("No es palindromo")

lista.mostrarLista()