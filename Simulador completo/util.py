def borde_por_rayo(x1, y1, xc, yc, x_min, y_min, x_max, y_max):
    """Devuelve el punto donde el rayo (x1,y1)->(xc,yc) intersecta el borde del rectángulo destino."""
    dx = xc - x1
    dy = yc - y1
    candidatos = []

    if dx != 0:
        t = (x_min - x1) / dx
        y = y1 + t * dy
        if t > 0 and y_min <= y <= y_max:
            candidatos.append((t, x_min, y))
        t = (x_max - x1) / dx
        y = y1 + t * dy
        if t > 0 and y_min <= y <= y_max:
            candidatos.append((t, x_max, y))

    if dy != 0:
        t = (y_min - y1) / dy
        x = x1 + t * dx
        if t > 0 and x_min <= x <= x_max:
            candidatos.append((t, x, y_min))
        t = (y_max - y1) / dy
        x = x1 + t * dx
        if t > 0 and x_min <= x <= x_max:
            candidatos.append((t, x, y_max))

    if not candidatos:
        x = min(max(xc, x_min), x_max)
        y = min(max(yc, y_min), y_max)
        return x, y

    _, x2, y2 = min(candidatos, key=lambda k: k[0])
    return x2, y2
