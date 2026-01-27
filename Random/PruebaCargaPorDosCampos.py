import csv as c
from datetime import datetime as dt

class Nodo:
    def __init__(self,regimen,fecha,nombre):
        self.reg = regimen 
        self.fecha = fecha 
        self.nombre = nombre 
        self.prox = None 

class Lista:
    def __init__(self):
        self.prim = None 

    def cargaOrdenada(self,regimen,fecha,nombre):
        q = Nodo(regimen,fecha,nombre)

        if self.prim is None:
            q.prox = self.prim 
            self.prim = q 
        else:
            ant = None
            p = self.prim 
            while p is not None and ((q.reg > p.reg) or ((q.reg == p.reg) and (q.fecha > p.fecha))):
                ant = p 
                p = p.prox 
            
            if p == self.prim:
                q.prox = p 
                self.prim = q 
            else:
                ant.prox = q 
                q.prox = p 

    def cargar(self,nombreArchivo):
        try:
            with open(nombreArchivo, mode='r',encoding="utf-8") as arch:
                lector = c.reader(arch)

                for fila in lector:
                    if not fila:
                        continue

                    regimen = fila[0]
                    fecha = dt.strptime(fila[1], "%d-%m-%Y") # Para transformar el string de la fecha en un objeto "Fecha"
                    nombre = fila[2]

                    self.cargaOrdenada(regimen,fecha,nombre)

        except FileNotFoundError:
            print("Error al leer el archivo")

    def mostrarNodo(self):
        p = self.prim 
        while p is not None:
            fecha_legible = p.fecha.strftime("%d-%m-%Y") # Para borrar el "00:00:00"
            print(f"Regimen: {p.reg} | Fecha: {fecha_legible} | Nombre: {p.nombre}")
            p = p.prox 
            
def long(A: Lista) -> int:
    if A.prox is None or A.prox == A:
        return 1 
    else: 
        return 1 + long(A.prox)
    
simple = Lista()
simple.cargar(r"Random/Personas.csv")

simple.mostrarNodo()

print(f"La longitud de la lista es de {long(simple.prim)} nodos")
