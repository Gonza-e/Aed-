import csv as c
from datetime import datetime as dt

def conv(x: str) -> int:
    if x == "A": return 1 
    elif x == "B": return 2 
    elif x == "C": return 3 
    elif x == "D": return 4 

def mostrar(x: int) -> str:
    if x == 1: return "A"
    elif x == 2: return "B"
    elif x == 3: return "C"
    elif x == 4: return "D"

class Regionales:
    def __init__(self,cantAlumnos,cantMaterias):
        self.cantAlum = cantAlumnos
        self.cantMat = cantMaterias

cant = [Regionales(0,0),Regionales(0,0),Regionales(0,0),Regionales(0,0),Regionales(0,0),Regionales(0,0)]

class Nodo:
    def __init__(self,regional,fecha,nombre,materias):
        self.reg = regional 
        self.fecha = fecha 
        self.nombre = nombre 
        self.materias = materias
        self.prox = None 

class Lista:
    def __init__(self):
        self.prim = None 

    def cargaOrdenada(self,regional,fecha,nombre,materias):
        q = Nodo(regional,fecha,nombre,materias)
        num = conv(q.reg)
        cant[num].cantAlum += 1
        cant[num].cantMat += q.materias

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

                    regional = fila[0]
                    fecha = dt.strptime(fila[1], "%d-%m-%Y") # Para transformar el string de la fecha en un objeto "Fecha"
                    nombre = fila[2]
                    materias = int(fila[3])

                    self.cargaOrdenada(regional,fecha,nombre,materias)
        except FileNotFoundError:
            print("Error al leer el archivo")

    def mostrarNodo(self):
        p = self.prim 
        while p is not None:
            fecha_legible = p.fecha.strftime("%d-%m-%Y") # Para borrar el "00:00:00"
            print(f"Regional: {p.reg} | Fecha: {fecha_legible} | Nombre: {p.nombre}")

            if (p.prox is None) or (p.prox).reg != p.reg:
                print(f"Regional: {p.reg} | Cantidad de alumnos: {cant[conv(p.reg)].cantAlum} | Cantidad de materias cursadas: {cant[conv(p.reg)].cantMat}")
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
