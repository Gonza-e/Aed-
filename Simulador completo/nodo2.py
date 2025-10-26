import tkinter as tk
from util import borde_por_rayo

class Nodo2:
    def __init__(self, canvas, x, y, valor):
        self.canvas = canvas
        self.x = x
        self.y = y
        self.valor = valor

        self.siguiente = None
        self.anterior = None
        self.flecha_siguiente = None
        self.flecha_anterior = None

        self.rect_dato = canvas.create_rectangle(x+20, y, x+80, y+30, fill="lightblue")
        self.rect_ptr1 = canvas.create_rectangle(x, y, x+20, y+30, fill="lightgreen")
        self.rect_ptr2 = canvas.create_rectangle(x+80, y, x+100, y+30, fill="lightgreen")
        self.texto = canvas.create_text(x+50, y+15, text=str(valor))

        self.items = [self.rect_dato, self.rect_ptr1, self.rect_ptr2, self.texto]

    def contiene(self, x, y):
        return self.x <= x <= self.x+100 and self.y <= y <= self.y+30

    def mover(self, dx, dy, objetos):
        self.x += dx
        self.y += dy
        for item in self.items:
            self.canvas.move(item, dx, dy)

        if self.siguiente:
            self.redibujar_flecha(self.siguiente, tipo="siguiente")
        if self.anterior:
            self.redibujar_flecha(self.anterior, tipo="anterior")

        for obj in objetos:
            if obj.siguiente == self:
                obj.redibujar_flecha(self, tipo="siguiente")
            if obj.anterior == self:
                obj.redibujar_flecha(self, tipo="anterior")

    def redibujar_flecha(self, otro, tipo):
        if tipo == "siguiente" and self.flecha_siguiente:
            self.canvas.delete(self.flecha_siguiente)
        if tipo == "anterior" and self.flecha_anterior:
            self.canvas.delete(self.flecha_anterior)

        if tipo == "siguiente":
            x1, y1 = self.x + 90, self.y + 15
        else:
            x1, y1 = self.x + 10, self.y + 15

        ancho_dest = 100 if hasattr(otro, "rect_dato") else 80
        x_min, y_min = otro.x, otro.y
        x_max, y_max = otro.x + ancho_dest, otro.y + 30
        xc, yc = otro.x + ancho_dest/2, otro.y + 15

        x2, y2 = borde_por_rayo(x1, y1, xc, yc, x_min, y_min, x_max, y_max)

        xm = (x1 + x2) / 2
        ym = (y1 + y2) / 2 - 40

        flecha = self.canvas.create_line(
            x1, y1, xm, ym, x2, y2,
            smooth=True, arrow=tk.LAST, width=2,
            fill="blue" if tipo=="siguiente" else "red"
        )

        if tipo == "siguiente":
            self.flecha_siguiente = flecha
            self.siguiente = otro
        else:
            self.flecha_anterior = flecha
            self.anterior = otro

    def eliminar_flecha(self, tipo=None):
        if tipo in (None, "siguiente"):
            if self.flecha_siguiente:
                self.canvas.delete(self.flecha_siguiente)
                self.flecha_siguiente = None
                self.siguiente = None
        if tipo in (None, "anterior"):
            if self.flecha_anterior:
                self.canvas.delete(self.flecha_anterior)
                self.flecha_anterior = None
                self.anterior = None
