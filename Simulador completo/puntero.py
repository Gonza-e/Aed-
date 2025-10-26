import tkinter as tk
from nodo import borde_por_rayo



class Puntero:
    def __init__(self, canvas, x, y, nombre):
        self.canvas = canvas
        self.x = x
        self.y = y
        self.nombre = nombre
        self.siguiente = None
        self.flecha = None

        self.rect = canvas.create_rectangle(x, y, x+80, y+30, fill="orange")
        self.texto = canvas.create_text(x+40, y+15, text=str(nombre))

        self.items = [self.rect, self.texto]

    def contiene(self, x, y):
        return self.x <= x <= self.x+80 and self.y <= y <= self.y+30

    def mover(self, dx, dy, objetos):
        self.x += dx
        self.y += dy
        for item in self.items:
            self.canvas.move(item, dx, dy)

        if self.flecha and self.siguiente:
            self.canvas.delete(self.flecha)
            self.conectar(self.siguiente)

        for obj in objetos:
            if obj.siguiente == self:
                if obj.flecha:
                    self.canvas.delete(obj.flecha)
                obj.conectar(self)

    def conectar(self, otro):
        if self.flecha:
            self.canvas.delete(self.flecha)
        self.siguiente = otro

        # Origen: centro del lado derecho del puntero
        x1, y1 = self.x + 80, self.y + 15

        # Rect destino (ancho 100 si es Nodo, 80 si es Puntero)
        ancho_dest = 100 if hasattr(otro, "rect_dato") else 80
        x_min, y_min = otro.x, otro.y
        x_max, y_max = otro.x + ancho_dest, otro.y + 30

        # Centro geométrico del destino
        xc, yc = otro.x + ancho_dest / 2, otro.y + 15

        # Punto de entrada correcto en el borde
        x2, y2 = borde_por_rayo(x1, y1, xc, yc, x_min, y_min, x_max, y_max)

        # Curva suave
        xm = (x1 + x2) / 2
        ym = (y1 + y2) / 2 - 40

        self.flecha = self.canvas.create_line(
            x1, y1, xm, ym, x2, y2,
            smooth=True, arrow=tk.LAST, width=2
        )


    def eliminar_flecha(self):
        if self.flecha:
            self.canvas.delete(self.flecha)
            self.flecha = None
            self.siguiente = None
