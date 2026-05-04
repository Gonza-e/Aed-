Accion ej1 es 
	Ambiente 
		jugador = registro 
			nombre: AN(30)
			estado: ("ACTIVO","ELIMINADO")
			b1: puntero a barcos1
			b2: puntero a barcos2 
			b4: puntero a barcos4
			disparos: N(2)
			prox: puntero a jugador 
		Freg 
		prim,p,q,x: puntero a jugador 

		tablero: arreglo[1...12,1...12] de entero

		coordenadas = registro 
			x,y: N(2)
		Freg  

		barcos1 = registro 
			imp: N(1)
			estado: ("HUNDIDO","UTIL")
			h1: coordenadas 
			prox: puntero a barcos1 
		Freg 
		prim1,p1,q1: puntero a barcos1

		barcos2 = registro 
			imp: N(1)
			estado: ("HUNDIDO","UTIL")
			h2: arreglo[1...2] de coordenadas
			prox: puntero a barcos2
		Freg 
		prim2,p2,q2: puntero a barcos2

		barcos4 = registro 
			imp: N(1)
			estado: ("HUNDIDO","UTIL")
			h4: arreglo[1...4] de coordenadas 
			prox: puntero a barcos4 
		Freg 
		prim4,p4,q4: puntero a barcos4

		Procedimiento cargaApilada(primero,nuevo)
			nuevo.prox:= primero 
			primero:= nuevo 
		FP 

		Procedimiento colocar1(jug: puntero a jugador) 
			Ambiente 
				contBarcos1: entero 
			Proceso 
				p1:= *q.barcos1
				contBarcos1:= 3
				Repetir 
					Esc("Ingrese las coordenadas X e Y del barco")
					Leer(cX); Leer(cY)

					si cX <= 10 y cY <= 10 entonces 
						Nuevo(q1); *q1.h1.x:= cX; *q1.h1.y:= cY; *q1.imp:= 1; *q1.estado:= "UTIL"
						contBarcos1:= contBarcos1-1 
						*q1.prox:= p1
						p1:= q1
					Fsi  
				Hasta (contBarcos1 = 0)
				*q.barcos1:= p1 
			FProceso 
		FP 

		Procedimiento colocar2(jug: puntero a jugador) 
			Ambiente 
				controlador,indice,contBarcos2: entero 
			Proceso 
				p2:= *q.barcos2
				contBarcos2:= 2; indice:= 0
				Repetir 
					controlador:= 2
					Nuevo(q2); *q2.imp:= 2; *q2.estado:= "UTIL"
					Repetir 
						Esc("Ingrese las coordenadas X e Y del barco")
						Leer(cX); Leer(cY)

						indice:= 3-controlador
						si ((cX <= 10) y (cY <= 10)) o ((indice = 2) y (((*q2.h2[2].x - 1 = cX) o (*q2.h2[2].x + 1 = cX)) y ((*q2.h2[2].y - 1 = cY) o (*q2.h2[2].y + 1 = cY)) y (cX <= 10 y cY))) entonces
							*q2.h[indice].x:= cX; *q2.h[indice].y:= cY 
							controlador:= controlador-1
						Fsi 
					Hasta (controlador = 0)
					*q1.prox:= p1 
					p1:= q1
					contBarcos2:= contBarcos2-1  
				Hasta (contBarcos2 = 0)
				*q.barcos2:= p2
			FProceso 
		FP 

		Procedimiento colocar4(jug: puntero a jugador) 
			Ambiente 
				controlador1,indice2: entero 
			Proceso 
				indice2:= 0
				controlador:= 4
				Nuevo(q4); *q4.imp:= 4; *q4.estado:= "UTIL"
				Repetir 
					Esc("Ingrese las coordenadas X e Y del barco")
					Leer(cX); Leer(cY)

					indice2:= 5-controlador
					si ((cX <= 10) y (cY <= 10)) o ((indice2 >= 2) y (((*q4.h3[indice2-1].x - 1 = cX) o (*q4.h3[indice2-1].x + 1 = cX)) y ((*q4.h3[indice2-1].y - 1 = cY) o (*q4.h3[indice2-1].y + 1 = cY)) y (cX <= 10 y cY))) entonces
						*q2.h[indice2].x:= cX; *q2.h[indice2].y:= cY 
						controlador1:= controlador1-1
					Fsi 
				Hasta (controlador1 = 0)
				*q.barcos4:= q4 
			FProceso 
		FP 

		Procedimiento initJugadores()
			Ambiente 
				nombre: AN 
			Proceso 
				Nuevo(q)
				Esc("Ingrese el nombre del jugador")
				Leer(nombre); *q.nombre:= nombre 

				colocar1(q); colocar2(q); colocar4()
				*q.disparos:= 25 
				cargaApilada(prim,q)
			FProceso 
		FP 

		Procedimiento jugar()
			Ambiente 
				cX1, cY1: entero 
			Proceso 
				p:= prim; x:= *p.prox 
				Mientras p <> nil hacer 
					Mientras x <> nil hacer 
						Esc("Jugador",*p.nombre,"ingrese coordenadas:")
						Leer(cX1); Leer(cY1)

						comprobar(cX1,cY1)

						Esc("Jugador",*x.nombre,"ingrese coordenadas:")
						Leer(cX1); Leer(cY1)

						comprobar(cX1,cY1)
						x:= *x.prox 
					FM 
					p:= *p.prox 
				FM 



			
				



