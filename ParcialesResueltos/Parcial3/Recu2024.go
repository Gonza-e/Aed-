Accion recu2024 (prim: puntero a nodo) es 
 Ambiente 
	nodo = registro 
		legajo: N(5)
		apeynom: AN(30)
		comision: (1...6)
		cant_instancias_aprobadas: N(2)
		notas: arreglo[1...11] de enteros 
		prox: puntero a nodo 
	Freg
	p,q,ant: puntero a nodo 

	nodo2 = registro 
		legajo: N(5)
		apeynom: AN(30)
		estado: ("Regular","Promocion")
		nota_prom: N(2,1)
		ant, prox: puntero a nodo2 
	Freg
	prim2,p2,q2,ult2: puntero a nodo2
	
	regulares: arreglo[1...6] de enteros 
	prom, i: entero

	funcion promedio(A: arreglo[1...11] de enteros, ind: entero): real 
		si i = 12 entonces 
			promedio:= A[ind]
		sino 
			si A[ind] >= 6
				promedio:= (A[ind]) + promedio(A,ind+1)
			fsi 
		fsi 
	FFuncion 
	
	procedimiento cargaOrdenada()
		si prim2 = nil entonces 
			*q2.prox:= nil 
			*q2.ant:= nil 
			prim2:= q 
			ult2:= q 
		sino 
			p2:= prim2 
			mientras p2 <> nil y (*q.nota_prom > *p.nota_prom)
				p2:= *p2.prox 
			FM 

			si p2 = prim2 entonces 
				*q2.prox:= p2
				*q2.ant:= nil 
				*p2.ant:= q2 
				prim2:= q2 
			sino 
				si p2 = ult2 entonces 
					*q2.prox:= *ult2.prox
					*q2.ant:= ult 
					ult2.prox:= q2 
					ult2:= q2 
				sino 
					*(*p2.ant).prox:= q2 
					*q2.ant:= *p2.ant 
					*q2.prox:= p2 
					*p2.ant:= q2 
				fsi 
			fsi 
		fsi 
	fprocedimiento 


	Proceso 
		prim2:= nil 
		Para i:= 1 hasta 6 hacer 
			regulares[i]:= 0 
		FPara 

		p:= prim 
		mientras p <> nil hacer 
			si *p.cant_instancias_aprobadas >= 5 entonces 
				Nuevo(q2); *q2.legajo:= *p.legajo; *q2.apeynom:= *p.apeynom
				segun *p.cant_instancias_aprobadas hacer 
					5: *q2.estado:= "Regular"; regulares[*p.comision]:= regulares[*p.comision] + 1 
					6: *q2.estado:= "Promocion"
				fsegun

				prom:= promedio(*p.notas,11) / *p.cant_instancias_aprobadas
				*q2.nota_prom:= prom 
				cargaOrdenada()
			fsi 
			p:= *p.prox 
		FM 

		p2:= ult2 
		Para i:= 1 hasta 10 hacer
			Esc("Legajo: ",*p2.legajo,"Nombre y apellido: ",*p2.apeynom,"Nota promedio: ",*p2.nota_prom)
			p2:= *p2.ant
		FPara 
		
		Para i:= 1 hasta 6 hacer 
			Esc("Comision: ",i,"Cantidad de alumnos regulares: ",regulares[i])
		FPara 
	FProceso 
FAccion



Accion recu2024_ejericio2 (prim, ult: puntero a nodo) es 
 Ambiente
	nodo = registro
		legajo: N(5)
		apeynom: AN(30)
		estado: ("Regular","Promocion")
		nota_prom: N(2,1)
		prox,ant: puntero a nodo 
	Freg
	p: puntero a nodo 

	FORMATO = registro 
		año: N(4)
		nro_acta: N(5)
	Freg
	reg: FORMATO
	
	procedimiento eliminar()
		si prim = nil entonces 
			Esc("La lista está vacia")
		sino 
			si p = prim entonces 
				prim:= *prim.prox 
				*prim.ant:= nil 
			sino 
				si p = ult entonces 
					ult:= *ult.ant 
					*ult.prox:= nil 
				sino 
					*(*p.ant).prox:= *p.prox 
					*(*p.prox).ant:= *p.ant 
				fsi 
			fsi 
		fsi 

		e:= p
		p:= *p.prox
		disponer(e)
	fprocedimiento

	cont: entero 

	Proceso
		cont:= 0
		p:= prim 
		reg:= nil 
		mientras p <> nil hacer 
			si *p.estado = "Regular" entonces 
				eliminar()
			sino 
				p:= *p.prox 
			fsi 
		FM

		reg:= genNroActa()
		Esc("Nro acta: ",reg.año,"/",reg.nro_acta)
		Esc("Legajo | Apellido y Nombre | Calificacion final")
		p:= prim
		mientras p <> nil hacer 
			Esc(*p.legajo,"|",*p.apeynom,"|",*p.nota_prom)
			p:= *p.prox 
			cont:= cont + 1
		FM

		Esc("Total de alumnos: ", cont)

	FProceso 
FAccion

			