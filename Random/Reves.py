class Nodo:
    def __init__(self,let):
        self.letra = let 
        self.prox = None 
    
class Lista:
    def __init__(self):
        self.prim = None 
    
    def cargaEncolada(self,let):
        q = Nodo(let)
        p = self.prim
        ant = None
        while p is not None:
            ant = p 
            p = p.prox 
        
        if self.prim is None:
            q.prox = self.prim 
            self.prim = q
        else: 
            q.prox = ant.prox
            ant.prox = q 
    
    def reves(self,nodo = None):
        if nodo is not None:
            self.reves(nodo.prox)
            print(f"{nodo.letra}")
        
palabra = 'venezuela'
lista = Lista()

for letra in palabra:
    lista.cargaEncolada(letra)


p = lista.prim
lista.reves(p)